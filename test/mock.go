package test

import (
	"github.com/cantillo16/bia_energy/src/models"
	"time"
)

type MockConsumptionRepository struct {
	GetMonthlyConsumptionFunc func(meterIDs []string, startDate, endDate time.Time) ([]models.Consumption, error)
	GetWeeklyConsumptionFunc  func(meterIDs []string, startDate, endDate time.Time) ([]models.Consumption, error)
	GetConsumptionDailyFunc   func(meterIDs []string, startDate, endDate time.Time) ([]models.Consumption, error)
	GetLastConsumptionFunc    func() ([]models.Consumption, error)
}

func (m *MockConsumptionRepository) GetMonthlyConsumption(
	meterIDs []string, startDate, endDate time.Time) (
	[]models.Consumption, error) {
	return m.GetMonthlyConsumptionFunc(meterIDs, startDate, endDate)
}

func (m *MockConsumptionRepository) GetConsumptionWeekly(meterIDs []string, startDate, endDate time.Time) (
	[]models.Consumption, error) {
	return m.GetWeeklyConsumptionFunc(meterIDs, startDate, endDate)
}

func (m *MockConsumptionRepository) GetConsumptionDaily(meterIDs []string, startDate, endDate time.Time) (
	[]models.Consumption, error) {
	return m.GetConsumptionDailyFunc(meterIDs, startDate, endDate)
}

func (m *MockConsumptionRepository) GetLastConsumption() ([]models.Consumption, error) {
	return m.GetLastConsumptionFunc()
}

type MockConsumptionService struct {
	GetConsumptionMonthlyFunc func(meterIDs []string, startDate, endDate time.Time) (
		[]models.Consumption, error)
	GetConsumptionWeeklyFunc func(meterIDs []string, startDate, endDate time.Time) (
		[]models.Consumption, error)
	GetConsumptionDailyFunc func(meterIDs []string, startDate, endDate time.Time) (
		[]models.Consumption, error)
	GetLastConsumptionFunc func() ([]models.Consumption, error)
}

func (m *MockConsumptionService) GetConsumptionMonthly(meterIDs []string, startDate, endDate time.Time) (
	[]models.Consumption, error) {
	return m.GetConsumptionMonthlyFunc(meterIDs, startDate, endDate)
}

func (m *MockConsumptionService) GetConsumptionWeekly(meterIDs []string, startDate, endDate time.Time) (
	[]models.Consumption, error) {
	return m.GetConsumptionWeeklyFunc(meterIDs, startDate, endDate)
}

func (m *MockConsumptionService) GetConsumptionDaily(meterIDs []string, startDate, endDate time.Time) (
	[]models.Consumption, error) {
	return m.GetConsumptionDailyFunc(meterIDs, startDate, endDate)
}

func (m *MockConsumptionService) GetLastConsumption() ([]models.Consumption, error) {
	return m.GetLastConsumptionFunc()
}
