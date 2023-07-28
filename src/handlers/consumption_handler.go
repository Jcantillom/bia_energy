package handlers

import (
	"encoding/json"
	"github.com/cantillo16/bia_energy/src/models"
	"github.com/cantillo16/bia_energy/src/services"
	"net/http"
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
		var dataGraph []models.Consumption

		switch kindPeriod {
		case "monthly":
			dataGraph, _ = service.GetConsumption(meterIDs, startDate, endDate)

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
			dataGraph, _ = service.GetConsumption(meterIDs, startDate, endDate)
			for i := 0; i < int(endDate.Sub(startDate).Hours()/24)+1; i++ {
				period = append(period, startDate.AddDate(0, 0, i).Format("Jan 2"))
			}

		default:
			http.Error(w, "El parámetro kind_period no es válido", http.StatusBadRequest)
			return
		}

		response := struct {
			Period    []string             `json:"period"`
			DataGraph []models.Consumption `json:"data_graph"`
		}{
			Period:    period,
			DataGraph: dataGraph,
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
