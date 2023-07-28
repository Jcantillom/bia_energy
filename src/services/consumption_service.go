package services

import (
	"github.com/cantillo16/bia_energy/src/models"
	"github.com/cantillo16/bia_energy/src/repositories"
	"time"
)

type ConsumptionService interface {
	GetConsumption(meterIDs []int, startDate, endDate time.Time) ([]models.Consumption, error)
	GetLastConsumption() ([]models.Consumption, error)
}

type ConsumptionServiceImpl struct {
	repo repositories.ConsumptionRepository
}

func NewConsumptionService(repo repositories.ConsumptionRepository) *ConsumptionServiceImpl {
	return &ConsumptionServiceImpl{repo: repo}
}

func (c *ConsumptionServiceImpl) GetConsumption(meterIDs []int, startDate, endDate time.Time) (
	[]models.Consumption, error) {
	return c.repo.GetConsumption(meterIDs, startDate, endDate)
}

func (c *ConsumptionServiceImpl) GetLastConsumption() ([]models.Consumption, error) {
	return c.repo.GetLastConsumption()
}
