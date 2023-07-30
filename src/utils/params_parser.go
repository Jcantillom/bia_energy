package utils

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func ParseRequestParams(r *http.Request) (time.Time, time.Time, string) {
	queryParams := r.URL.Query()
	startDate := parseDate(queryParams.Get("start_date"))
	endDate := parseDate(queryParams.Get("end_date"))

	// Convertir las fechas a la zona horaria UTC
	startDate = startDate.UTC()
	endDate = endDate.UTC()

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

	return parsedDate
}

func SetStartAndEndOfMonth(startDate, endDate time.Time) (time.Time, time.Time) {
	// Convertir las fechas a UTC
	startDate = startDate.UTC()
	endDate = endDate.UTC()

	// Establecer la hora de inicio del día (00:00:00)
	if startDate.Day() != 1 {
		startDate = startDate.AddDate(0, 0, -startDate.Day()+1)
	}
	startDate = startDate.Add(time.Hour*time.Duration(0) +
		time.Minute*time.Duration(0) + time.Second*time.Duration(0))

	// Establecer la hora de fin del día (23:59:59)
	if endDate.Day() != endDate.AddDate(0, 1, -endDate.Day()).Day() {
		endDate = endDate.AddDate(0, 1, -endDate.Day())
	}
	endDate = endDate.Add(time.Hour*23 + time.Minute*59 + time.Second*59)

	return startDate, endDate
}

func SetStartAndEndWeek(startDate, endDate time.Time) (time.Time, time.Time) {
	startDate = startDate.UTC()
	endDate = endDate.UTC()

	// Establecer la hora de inicio del día (00:00:00)
	startDate = startDate.Add(time.Hour*time.Duration(0) +
		time.Minute*time.Duration(0) + time.Second*time.Duration(0))

	// Establecer la hora de fin del día (23:59:59)
	endDate = endDate.Add(time.Hour*23 + time.Minute*59 + time.Second*59)

	// Crear una copia de startDate y avanzar 7 días para calcular el siguiente intervalo
	nextStartDate := startDate
	nextStartDate = nextStartDate.AddDate(0, 0, 7)

	var periods []string

	// Calcular los intervalos de 7 días hasta que lleguemos a la endDate
	for nextStartDate.Before(endDate) {
		period := fmt.Sprintf("%s - %s", startDate.Format("JAN 2"), nextStartDate.Add(-time.Second).Format("JAN 2"))
		periods = append(periods, period)

		// Avanzar 7 días para el próximo intervalo
		startDate = startDate.AddDate(0, 0, 7)
		nextStartDate = nextStartDate.AddDate(0, 0, 7)
	}

	// Agregar el último período desde la última startDate hasta endDate
	lastPeriod := fmt.Sprintf("%s - %s", startDate.Format("JAN 2"), endDate.Format("JAN 2"))
	periods = append(periods, lastPeriod)

	// Retornar las fechas ajustadas y los períodos de 7 días
	return startDate, endDate
}

func SetStartAndEndDay(startDate, endDate time.Time) (time.Time, time.Time) {
	// Convertir las fechas a UTC
	startDate = startDate.UTC()
	endDate = endDate.UTC()

	// Establecer la hora de inicio del día (00:00:00)
	startDate = startDate.Add(time.Hour*time.Duration(0) +
		time.Minute*time.Duration(0) + time.Second*time.Duration(0))

	// Establecer la hora de fin del día (23:59:59)
	endDate = endDate.Add(time.Hour*23 + time.Minute*59 + time.Second*59)

	return startDate, endDate
}
