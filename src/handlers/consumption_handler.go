package handlers

import (
	"encoding/json"
	"github.com/cantillo16/bia_energy/src/models"
	"github.com/cantillo16/bia_energy/src/services"
	"github.com/cantillo16/bia_energy/src/utils"
	"net/http"
)

type (
	Controller func(w http.ResponseWriter, r *http.Request)
	Handler    struct {
		GetConsumption Controller
		GetLast        Controller
	}
)

func NewHandler(service services.ConsumptionService) *Handler {
	return &Handler{
		GetConsumption: getConsumptionHandler(service),
		GetLast:        getLastHandler(service),
	}
}

func getConsumptionHandler(service services.ConsumptionService) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		startDate, endDate, kindPeriod := utils.ParseRequestParams(r)
		meterIDs := utils.ParseMeterIDs(r.URL.Query().Get("meters_ids"))

		if startDate.After(endDate) {
			http.Error(w, "La fecha inicial no puede ser mayor que la fecha final", http.StatusBadRequest)
			return
		}

		var dataGraph []models.Consumption
		var err error

		switch kindPeriod {
		case "monthly":
			dataGraph, err = service.GetMonthlyConsumption(meterIDs, startDate, endDate)
		case "daily":
			dataGraph, err = service.GetDailyConsumption(meterIDs, startDate, endDate)
		case "weekly":
			dataGraph, err = service.GetConsumptionWeekly(meterIDs, startDate, endDate)
		default:
			http.Error(w, "El valor del campo 'kind_period' es inválido", http.StatusBadRequest)
			return
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		groups := utils.GroupConsumptions(dataGraph, kindPeriod)

		// Generamos el período según el tipo de período seleccionado
		period := utils.GeneratePeriod(kindPeriod, startDate, endDate)

		// Agrupamos los valores por mes en los arreglos correspondientes
		var dataGraphActiveEnergy []float64
		var dataGraphReactiveEnergy []float64
		var dataGraphSolar []float64

		// Calcular el acumulado de los valores por mes y agregarlos a los arreglos
		for _, group := range groups {
			dataGraphActiveEnergy = append(dataGraphActiveEnergy, utils.SumarArreglo(group.ActiveEnergy))
			dataGraphReactiveEnergy = append(dataGraphReactiveEnergy, utils.SumarArreglo(group.ReactiveEnergy))
			dataGraphSolar = append(dataGraphSolar, utils.SumarArreglo(group.Solar))
		}

		// Construimos la respuesta JSON con la estructura deseada
		response := struct {
			Period    []string `json:"period"`
			DataGraph []struct {
				MeterID        string    `json:"MeterID"`
				ActiveEnergy   []float64 `json:"ActiveEnergy"`
				ReactiveEnergy []float64 `json:"ReactiveEnergy"`
				Solar          []float64 `json:"Solar"`
			} `json:"data_graph"`
		}{
			Period: period,
			DataGraph: []struct {
				MeterID        string    `json:"MeterID"`
				ActiveEnergy   []float64 `json:"ActiveEnergy"`
				ReactiveEnergy []float64 `json:"ReactiveEnergy"`
				Solar          []float64 `json:"Solar"`
			}{
				{
					MeterID:        groups[0].MeterID,
					ActiveEnergy:   dataGraphActiveEnergy,
					ReactiveEnergy: dataGraphReactiveEnergy,
					Solar:          dataGraphSolar,
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func getLastHandler(service services.ConsumptionService) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		lastConsumption, _ := service.GetLastConsumption()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(lastConsumption)
	}
}
