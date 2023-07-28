package test

import (
	"github.com/cantillo16/bia_energy/src/models"
	"time"
)

type MockConsumptionRepository struct {
	GetConsumptionFunc     func(meterIDs []int, startDate, endDate time.Time) ([]models.Consumption, error)
	GetLastConsumptionFunc func() ([]models.Consumption, error)
}

func (m *MockConsumptionRepository) GetConsumption(meterIDs []int, startDate, endDate time.Time) (
	[]models.Consumption, error) {
	return m.GetConsumptionFunc(meterIDs, startDate, endDate)
}

func (m *MockConsumptionRepository) GetLastConsumption() ([]models.Consumption, error) {
	return m.GetLastConsumptionFunc()
}

type MockConsumptionService struct {
	GetConsumptionFunc     func(meterIDs []int, startDate, endDate time.Time) ([]models.Consumption, error)
	GetLastConsumptionFunc func() ([]models.Consumption, error)
}

func (m *MockConsumptionService) GetConsumption(meterIDs []int, startDate, endDate time.Time) (
	[]models.Consumption, error) {
	return m.GetConsumptionFunc(meterIDs, startDate, endDate)
}

func (m *MockConsumptionService) GetLastConsumption() ([]models.Consumption, error) {
	return m.GetLastConsumptionFunc()
}
