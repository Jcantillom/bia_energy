package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/cantillo16/bia_energy/src/models"
	"github.com/cantillo16/bia_energy/src/services"
	"github.com/cantillo16/bia_energy/src/utils"
	"net/http"
	"sort"
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
		fmt.Println("startDate: ", startDate)
		fmt.Println("endDate: ", endDate)
		var period []string
		var dataGraphGroup []ConsumptionGroup

		var dataGraph []models.Consumption
		var err error

		switch kindPeriod {
		case "monthly":
			dataGraph, err = service.GetMonthlyConsumption(meterIDs, startDate, endDate)
		default:
			dataGraph, err = service.GetConsumption(meterIDs, startDate, endDate)
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Creamos una función para agrupar los datos por el tipo de periodo (monthly, weekly, daily)
		groupFunc := func(consumption models.Consumption) string {
			switch kindPeriod {
			case "monthly":
				return consumption.Date.Format("Jan 2006")
			case "weekly":
				return consumption.Date.Format("Jan 2 2006")
			case "daily":
				return consumption.Date.Format("02 Jan 2006")
			default:
				return ""
			}
		}

		// Creamos un mapa para agrupar los datos
		groupMap := make(map[string]*ConsumptionGroup)

		for _, consumption := range dataGraph {
			groupKey := groupFunc(consumption)
			if groupKey == "" {
				http.Error(w, "El parámetro kind_period no es válido", http.StatusBadRequest)
				return
			}

			group, ok := groupMap[groupKey]
			if !ok {
				groupMap[groupKey] = &ConsumptionGroup{
					MeterID:        consumption.MeterID,
					ActiveEnergy:   []float64{consumption.ActiveEnergy},
					ReactiveEnergy: []float64{consumption.ReactiveEnergy},
					Solar:          []float64{consumption.Solar},
					Date:           []string{consumption.Date.Format("2006-01-02")},
				}
			} else {
				group.ActiveEnergy = append(group.ActiveEnergy, consumption.ActiveEnergy)
				group.ReactiveEnergy = append(group.ReactiveEnergy, consumption.ReactiveEnergy)
				group.Solar = append(group.Solar, consumption.Solar)
				group.Date = append(group.Date, consumption.Date.Format("2006-01-02"))
			}
		}

		// Obtenemos la lista final de ConsumptionGroup y formateamos el periodo
		for _, group := range groupMap {
			dataGraphGroup = append(dataGraphGroup, *group)
		}

		// Ordenamos los datos por fecha antes de generar el periodo
		sort.SliceStable(dataGraphGroup, func(i, j int) bool {
			return dataGraphGroup[i].Date[0] < dataGraphGroup[j].Date[0]
		})

		// Generamos el periodo según el tipo de periodo seleccionado
		switch kindPeriod {
		case "monthly":
			if startDate.Day() == 1 {
				period = append(period, startDate.Format("Jan 2006"))
			}
			for startDate.Before(endDate) {
				startDate = startDate.AddDate(0, 1, 0)
				if startDate.Day() == 1 && startDate.Before(endDate) {
					period = append(period, startDate.Format("Jan 2006"))
				}
			}
		case "weekly":
			dataGraph, _ = service.GetConsumption(meterIDs, startDate, endDate)
			for startDate.Before(endDate) {
				weekEndDate := startDate.AddDate(0, 0, 6)
				if weekEndDate.After(endDate) {
					weekEndDate = endDate
				}
				period = append(period, startDate.Format("Jan 2")+" - "+weekEndDate.Format("Jan 2"))
				startDate = weekEndDate.AddDate(0, 0, 1)
			}
		case "daily":
			if startDate.Hour() == 0 {
				period = append(period, startDate.Format("02 Jan 2006"))
			}
			for startDate.Before(endDate) {
				startDate = startDate.AddDate(0, 0, 1)
				if startDate.Hour() == 0 && startDate.Before(endDate) {
					period = append(period, startDate.Format("02 Jan 2006"))
				}
			}
		default:
			http.Error(w, "El parámetro kind_period no es válido", http.StatusBadRequest)
			return
		}

		response := struct {
			Period    []string           `json:"period"`
			DataGraph []ConsumptionGroup `json:"data_graph"`
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
