package utils

import (
	"net/http"
	"strings"
	"time"
)

func ParseRequestParams(r *http.Request) (time.Time, time.Time, string) {
	queryParams := r.URL.Query()
	startDate := parseDate(queryParams.Get("start_date"))
	endDate := parseDate(queryParams.Get("end_date"))
	kindPeriod := queryParams.Get("kind_period")
	return startDate, endDate, kindPeriod
}

func ParseMeterIDs(meterIDs string) []string {
	if meterIDs == "" {
		return []string{}
	}
	return strings.Split(meterIDs, ",")
}

func parseDate(date string) time.Time {
	if date == "" {
		return time.Time{}
	}

	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Time{}
	}

	//Si la fecha no tiene hora, tomamos las 23:59:59 del día indicado
	//if !strings.Contains(date, "T") {
	//	parsedDate = parsedDate.Add(time.Hour*time.Duration(23) +
	//		time.Minute*time.Duration(59) +
	//		time.Second*time.Duration(59))
	//}

	return parsedDate
}

func SetStartAndEndOfMonth(startDate, endDate time.Time) (time.Time, time.Time) {
	// Establecer la hora de inicio del día (00:00:00)
	if startDate.Day() != 1 {
		startDate = startDate.AddDate(0, 0, -startDate.Day()+1)
	}
	startDate = startDate.Add(time.Hour*time.Duration(0) + time.Minute*time.Duration(0) + time.Second*time.Duration(0))
	//imprimir el type de startDate

	// Establecer la hora de fin del día (23:59:59)
	if endDate.Day() != endDate.AddDate(0, 1, -endDate.Day()).Day() {
		endDate = endDate.AddDate(0, 1, -endDate.Day())
	}
	endDate = endDate.Add(time.Hour*23 + time.Minute*59 + time.Second*59)

	return startDate, endDate
}
