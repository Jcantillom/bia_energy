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
			http.Error(w, "El valor del campo 'kind_period' es inválido",
				http.StatusBadRequest)
			return

		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		groupMap := utils.GroupConsumptions(dataGraph, kindPeriod)

		// Obtenemos la lista final de ConsumptionGroup y formateamos el período
		var dataGraphGroup []utils.ConsumptionGroup
		for _, group := range groupMap {
			dataGraphGroup = append(dataGraphGroup, *group)
		}

		// Realizar la suma de los acumulados en cada arreglo
		utils.SumarAcomulados(dataGraphGroup)

		// Generamos el período según el tipo de período seleccionado
		period := utils.GeneratePeriod(kindPeriod, startDate, endDate)

		response := struct {
			Period    []string                 `json:"period"`
			DataGraph []utils.ConsumptionGroup `json:"data_graph"`
		}{
			Period:    period,
			DataGraph: dataGraphGroup,
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
