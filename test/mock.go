package test

import (
	"github.com/cantillo16/bia_energy/src/models"
	"time"
)

type MockConsumptionRepository struct {
	GetConsumptionMonthlyFunc func(meterIDs []int, startDate, endDate time.Time) ([]models.Consumption, error)
	GetConsumptionDailyFunc   func(meterIDs []int, startDate, endDate time.Time) ([]models.Consumption, error)
	GetConsumptionWeeklyFunc  func(meterIDs []int, startDate, endDate time.Time) ([]models.Consumption, error)
}

func (m *MockConsumptionRepository) GetConsumptionMonthly(meterIDs []int, startDate, endDate time.Time) (
	[]models.Consumption, error) {
	return m.GetConsumptionMonthlyFunc(meterIDs, startDate, endDate)
}

func (m *MockConsumptionRepository) GetConsumptionDaily(meterIDs []int, startDate, endDate time.Time) (
	[]models.Consumption, error) {
	return m.GetConsumptionDailyFunc(meterIDs, startDate, endDate)
}

func (m *MockConsumptionRepository) GetConsumptionWeekly(meterIDs []int, startDate, endDate time.Time) (
	[]models.Consumption, error) {
	return m.GetConsumptionWeeklyFunc(meterIDs, startDate, endDate)
}

type MockConsumptionService struct {
	GetConsumptionMonthlyFunc func(meterIDs []int, startDate, endDate time.Time) ([]models.Consumption, error)
	GetConsumptionDailyFunc   func(meterIDs []int, startDate, endDate time.Time) ([]models.Consumption, error)
	GetConsumptionWeeklyFunc  func(meterIDs []int, startDate, endDate time.Time) ([]models.Consumption, error)
}

func (m *MockConsumptionService) GetConsumptionMonthly(meterIDs []int, startDate, endDate time.Time) (
	[]models.Consumption, error) {
	return m.GetConsumptionMonthlyFunc(meterIDs, startDate, endDate)
}

func (m *MockConsumptionService) GetConsumptionDaily(meterIDs []int, startDate, endDate time.Time) (
	[]models.Consumption, error) {
	return m.GetConsumptionDailyFunc(meterIDs, startDate, endDate)
}

func (m *MockConsumptionService) GetConsumptionWeekly(meterIDs []int, startDate, endDate time.Time) (
	[]models.Consumption, error) {
	return m.GetConsumptionWeeklyFunc(meterIDs, startDate, endDate)
}
