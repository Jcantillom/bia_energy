package repositories

import (
	"github.com/cantillo16/bia_energy/src/models"
	"github.com/cantillo16/bia_energy/src/utils"
	"gorm.io/gorm"
	"sort"
	"time"
)

type ConsumptionRepository interface {
	GetConsumption(meterIDs []string, startDate, endDate time.Time) ([]models.Consumption, error)
	GetLastConsumption() ([]models.Consumption, error)
	GetMonthlyConsumption(meterIDs []string, startDate, endDate time.Time) ([]models.Consumption, error)
}

type ConsumptionRepositoryImpl struct {
	db *gorm.DB
}

func NewConsumptionRepository(db *gorm.DB) *ConsumptionRepositoryImpl {
	return &ConsumptionRepositoryImpl{db: db}
}

func (c *ConsumptionRepositoryImpl) GetConsumption(meterIDs []string, startDate, endDate time.Time) (
	[]models.Consumption, error) {
	var consumptions []models.Consumption
	err := c.db.Where("meter_id IN (?) AND date BETWEEN ? AND ?",
		meterIDs, startDate, endDate).Find(&consumptions).Error
	if err != nil {
		return nil, err
	}
	return consumptions, nil
}

// obtener los ultimos 20 registros
func (c *ConsumptionRepositoryImpl) GetLastConsumption() ([]models.Consumption, error) {
	var consumptions []models.Consumption
	err := c.db.Order("date desc").Limit(20).Find(&consumptions).Error
	if err != nil {
		return nil, err
	}
	return consumptions, nil
}

func (c *ConsumptionRepositoryImpl) GetMonthlyConsumption(meterIDs []string, startDate, endDate time.Time) ([]models.Consumption, error) {
	var consumptions []models.Consumption

	startDate, endDate = utils.SetStartAndEndOfMonth(startDate, endDate)

	err := c.db.Where("meter_id IN (?) AND date BETWEEN ? AND ?",
		meterIDs, startDate, endDate).Find(&consumptions).Error
	if err != nil {
		return nil, err
	}

	// Ordenar las consumptions por fecha antes de devolverlas
	sort.SliceStable(consumptions, func(i, j int) bool {
		return consumptions[i].Date.Before(consumptions[j].Date)
	})

	return consumptions, nil
}
