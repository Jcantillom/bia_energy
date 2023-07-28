package handlers

import (
	"encoding/json"
	"github.com/cantillo16/bia_energy/src/models"
	"github.com/cantillo16/bia_energy/src/services"
	"net/http"
	"sort"
	"strconv"
	"time"
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
		meterIDs, startDate, endDate, kindPeriod := parseRequestParams(r)
		var period []string
		var dataGraphGroup []ConsumptionGroup

		dataGraph, _ := service.GetConsumption(meterIDs, startDate, endDate)

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

func parseRequestParams(r *http.Request) ([]int, time.Time, time.Time, string) {
	queryParams := r.URL.Query()
	meterIDs := parseMeterIDs(queryParams.Get("meters_ids"))
	startDate := parseDate(queryParams.Get("start_date"))
	endDate := parseDate(queryParams.Get("end_date"))
	kindPeriod := queryParams.Get("kind_period")
	return meterIDs, startDate, endDate, kindPeriod
}

func parseMeterIDs(meterIDs string) []int {
	if meterIDs == "" {
		return []int{}
	}
	var ids []int
	id, err := strconv.Atoi(meterIDs)
	if err != nil {
		http.Error(nil, "El parámetro meters_ids no es válido", http.StatusBadRequest)
		return []int{}
	}
	ids = append(ids, id)
	return ids
}

func parseDate(date string) time.Time {
	if date == "" {
		return time.Time{}
	}
	parsedDate, _ := time.Parse("2006-01-02", date)
	return parsedDate
}

func getLastHandler(service services.ConsumptionService) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		lastConsumption, _ := service.GetLastConsumption()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(lastConsumption)
	}
}
