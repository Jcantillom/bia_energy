package services

import (
	"github.com/cantillo16/bia_energy/src/models"
	"github.com/cantillo16/bia_energy/src/repositories"
	"time"
)

type ConsumptionService interface {
	GetConsumptionWeekly(meterIDs []string, startDate, endDate time.Time) ([]models.Consumption, error)
	GetMonthlyConsumption(meterIDs []string, startDate, endDate time.Time) ([]models.Consumption, error)
	GetDailyConsumption(meterIDs []string, startDate, endDate time.Time) ([]models.Consumption, error)
	GetLastConsumption() ([]models.Consumption, error)
}

type ConsumptionServiceImpl struct {
	repo repositories.ConsumptionRepository
}

func NewConsumptionService(repo repositories.ConsumptionRepository) *ConsumptionServiceImpl {
	return &ConsumptionServiceImpl{repo: repo}
}
func (c *ConsumptionServiceImpl) GetMonthlyConsumption(meterIDs []string, startDate, endDate time.Time) (
	[]models.Consumption, error) {
	return c.repo.GetMonthlyConsumption(meterIDs, startDate, endDate)
}
func (c *ConsumptionServiceImpl) GetConsumptionWeekly(meterIDs []string, startDate, endDate time.Time) (
	[]models.Consumption, error) {
	return c.repo.GetConsumptionWeekly(meterIDs, startDate, endDate)
}

func (c *ConsumptionServiceImpl) GetDailyConsumption(meterIDs []string, startDate, endDate time.Time) (
	[]models.Consumption, error) {
	return c.repo.GetConsumptionDaily(meterIDs, startDate, endDate)
}

func (c *ConsumptionServiceImpl) GetLastConsumption() ([]models.Consumption, error) {
	return c.repo.GetLastConsumption()
}
