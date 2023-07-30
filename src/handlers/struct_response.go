package handlers

type ConsumptionGroup struct {
	MeterID            string    `json:"meter_id"`
	ActiveEnergy       []float64 `json:"active_energy"`
	ReactiveEnergy     []float64 `json:"reactive_energy"`
	CapacitiveReactive []float64 `json:"capacitive_reactive"`
	Solar              []float64 `json:"solar"`
	Date               []string  `json:"date"`
}
