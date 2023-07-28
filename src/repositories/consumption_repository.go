package repositories

import (
	"github.com/cantillo16/bia_energy/src/models"
	"gorm.io/gorm"
	"time"
)

type ConsumptionRepository interface {
	GetConsumption(meterIDs []int, startDate, endDate time.Time) ([]models.Consumption, error)
	GetLastConsumption() ([]models.Consumption, error)
}

type ConsumptionRepositoryImpl struct {
	db *gorm.DB
}

func NewConsumptionRepository(db *gorm.DB) *ConsumptionRepositoryImpl {
	return &ConsumptionRepositoryImpl{db: db}
}

func (c *ConsumptionRepositoryImpl) GetConsumption(meterIDs []int, startDate, endDate time.Time) (
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
