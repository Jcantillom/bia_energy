package models

import (
	"time"
)

type Consumption struct {
	ID                 string    `json:"id" gorm:"type:char(36);not null;primary_key;unique_index"`
	MeterID            string    `json:"meter_id" gorm:"type:char;not null;index"`
	ActiveEnergy       float64   `json:"active_energy" gorm:"type:decimal(15,6);not null;index"`
	ReactiveEnergy     float64   `json:"reactive_energy" gorm:"type:decimal(15,6);not null;index"`
	CapacitiveReactive float64   `json:"capacitive_reactive" gorm:"type:decimal(15,6);not null;index"`
	Solar              float64   `json:"solar" gorm:"type:decimal(15,6);not null;index"`
	Date               time.Time `json:"date" gorm:"type:timestamp;not null;index"`
	CreateAt           time.Time `json:"create_at" gorm:"type:timestamp;null;index"`
	UpdateAt           time.Time `json:"update_at" gorm:"type:timestamp;null;index"`
}
