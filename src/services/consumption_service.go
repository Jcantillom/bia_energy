package services

import (
	"github.com/cantillo16/bia_energy/src/models"
	"github.com/cantillo16/bia_energy/src/repositories"
	"time"
)

type ConsumptionService interface {
	GetConsumptionMonthly(meterIDs []int, startDate, endDate time.Time) ([]models.Consumption, error)
	GetConsumptionDaily(meterIDs []int, startDate, endDate time.Time) ([]models.Consumption, error)
	GetConsumptionWeekly(meterIDs []int, startDate, endDate time.Time) ([]models.Consumption, error)
	GetLastConsumption() ([]models.Consumption, error)
}

type ConsumptionServiceImpl struct {
	repo repositories.ConsumptionRepository
}

func NewConsumptionService(repo repositories.ConsumptionRepository) *ConsumptionServiceImpl {
	return &ConsumptionServiceImpl{repo: repo}
}

func (c *ConsumptionServiceImpl) GetConsumptionMonthly(meterIDs []int, startDate, endDate time.Time) (
	[]models.Consumption, error) {
	return c.repo.GetConsumptionMonthly(meterIDs, startDate, endDate)
}

func (c *ConsumptionServiceImpl) GetConsumptionDaily(meterIDs []int, startDate, endDate time.Time) (
	[]models.Consumption, error) {
	return c.repo.GetConsumptionDaily(meterIDs, startDate, endDate)
}

func (c *ConsumptionServiceImpl) GetConsumptionWeekly(meterIDs []int, startDate, endDate time.Time) (
	[]models.Consumption, error) {
	return c.repo.GetConsumptionWeekly(meterIDs, startDate, endDate)
}

func (c *ConsumptionServiceImpl) GetLastConsumption() ([]models.Consumption, error) {
	return c.repo.GetLastConsumption()
}
