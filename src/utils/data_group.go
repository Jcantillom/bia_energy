package utils

import (
	"github.com/cantillo16/bia_energy/src/models"
	"time"
)

type ConsumptionGroup struct {
	MeterID        string    `json:"MeterID"`
	ActiveEnergy   []float64 `json:"ActiveEnergy"`
	ReactiveEnergy []float64 `json:"ReactiveEnergy"`
	Solar          []float64 `json:"Solar"`
}

func GroupConsumptions(dataGraph []models.Consumption, kindPeriod string) []*ConsumptionGroup {
	groupMap := make(map[string]*ConsumptionGroup)

	// Creamos una función para obtener la clave del grupo por el tipo de período (monthly, weekly, daily)
	getGroupKey := func(consumption models.Consumption) string {
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

	// Agrupamos los datos y sumamos los valores por mes
	for _, consumption := range dataGraph {
		groupKey := getGroupKey(consumption)
		if groupKey == "" {
			continue
		}

		group, ok := groupMap[groupKey]
		if !ok {
			groupMap[groupKey] = &ConsumptionGroup{
				MeterID:        consumption.MeterID,
				ActiveEnergy:   []float64{consumption.ActiveEnergy},
				ReactiveEnergy: []float64{consumption.ReactiveEnergy},
				Solar:          []float64{consumption.Solar},
			}
		} else {
			group.ActiveEnergy = append(group.ActiveEnergy, consumption.ActiveEnergy)
			group.ReactiveEnergy = append(group.ReactiveEnergy, consumption.ReactiveEnergy)
			group.Solar = append(group.Solar, consumption.Solar)
		}
	}

	// Convertir el mapa en una lista de objetos ConsumptionGroup
	var result []*ConsumptionGroup
	for _, group := range groupMap {
		result = append(result, group)
	}

	return result
}

// GeneratePeriod genera el período según el tipo de período seleccionado.
func GeneratePeriod(kindPeriod string, startDate, endDate time.Time) []string {
	var period []string

	// Función para verificar si una fecha debe agregarse al período
	shouldAddDate := func(date time.Time) bool {
		switch kindPeriod {
		case "monthly":
			return true // Siempre agregamos el mes completo
		case "weekly":
			// En este caso, siempre agregamos la fecha, ya que el período es semanal
			return true
		case "daily":
			return true
		default:
			return false
		}
	}

	for startDate.Before(endDate) {
		// Verificar si la fecha actual debe agregarse al período
		if shouldAddDate(startDate) {
			switch kindPeriod {
			case "monthly":
				period = append(period, startDate.Format("Jan 2006"))
			case "weekly":
				weekEndDate := startDate.AddDate(0, 0, 6)
				if weekEndDate.After(endDate) {
					weekEndDate = endDate
				}
				period = append(period, startDate.Format("Jan 2")+" - "+weekEndDate.Format("Jan 2"))
			case "daily":
				period = append(period, startDate.Format("02 Jan 2006"))
			}
		}

		// Avanzar al siguiente período
		switch kindPeriod {
		case "monthly":
			startDate = startDate.AddDate(0, 1, 0) // Avanzar un mes
		case "weekly":
			startDate = startDate.AddDate(0, 0, 7) // Avanzar una semana
		case "daily":
			startDate = startDate.AddDate(0, 0, 1) // Avanzar un día
		}
	}

	return period
}
