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

		var dataGraph []models.Consumption
		var err error

		switch kindPeriod {
		case "monthly":
			dataGraph, err = service.GetMonthlyConsumption(meterIDs, startDate, endDate)
		case "daily":
			dataGraph, err = service.GetDailyConsumption(meterIDs, startDate, endDate)
		case "weekly":
			dataGraph, err = service.GetConsumptionWeekly(meterIDs, startDate, endDate)
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		period := utils.GeneratePeriod(kindPeriod, startDate, endDate)
		dataGraphGroup := utils.GroupConsumptions(dataGraph, kindPeriod)
		// Verificar si dataGraphGroup está vacío, lo que significa que no se encontraron consumos para los medidores especificados
		if len(dataGraphGroup) == 0 {
			http.Error(w, "No se encontraron consumos para los medidores especificados",
				http.StatusNotFound)
			return
		}
		dataGraphActiveEnergy := utils.GetSumarizedData(dataGraphGroup, "ActiveEnergy")
		dataGraphReactiveEnergy := utils.GetSumarizedData(dataGraphGroup, "ReactiveEnergy")
		dataGraphSolar := utils.GetSumarizedData(dataGraphGroup, "Solar")

		response := buildResponse(period, meterIDs, dataGraphActiveEnergy,
			dataGraphReactiveEnergy, dataGraphSolar)

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

func buildResponse(
	period []string,
	meterIDs []string,
	activeEnergy []float64,
	reactiveEnergy []float64,
	solar []float64) interface{} {
	dataGraphGroup := make([]struct {
		MeterID        string    `json:"MeterID"`
		ActiveEnergy   []float64 `json:"ActiveEnergy"`
		ReactiveEnergy []float64 `json:"ReactiveEnergy"`
		Solar          []float64 `json:"Solar"`
	}, len(meterIDs))

	for i, meterID := range meterIDs {
		dataGraphGroup[i] = struct {
			MeterID        string    `json:"MeterID"`
			ActiveEnergy   []float64 `json:"ActiveEnergy"`
			ReactiveEnergy []float64 `json:"ReactiveEnergy"`
			Solar          []float64 `json:"Solar"`
		}{
			MeterID:        meterID,
			ActiveEnergy:   activeEnergy,
			ReactiveEnergy: reactiveEnergy,
			Solar:          solar,
		}
	}

	return struct {
		Period    []string `json:"period"`
		DataGraph []struct {
			MeterID        string    `json:"MeterID"`
			ActiveEnergy   []float64 `json:"ActiveEnergy"`
			ReactiveEnergy []float64 `json:"ReactiveEnergy"`
			Solar          []float64 `json:"Solar"`
		} `json:"data_graph"`
	}{
		Period:    period,
		DataGraph: dataGraphGroup,
	}
}
