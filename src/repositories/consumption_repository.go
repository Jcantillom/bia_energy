package repositories

import (
	"fmt"
	"github.com/cantillo16/bia_energy/src/models"
	"gorm.io/gorm"
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

	// Ajustar las fechas para obtener el rango de meses completo
	startDate = time.Date(startDate.Year(), startDate.Month(), 1, 0, 0, 0, 0, time.UTC)
	endDate = time.Date(endDate.Year(), endDate.Month()+1, 1, 0, 0, 0, 0, time.UTC).Add(-time.Nanosecond)

	// Si la fecha de inicio es antes del primer día del mes seleccionado, ajustarla al primer día
	if startDate.Before(time.Date(endDate.Year(), endDate.Month(), 1, 0, 0, 0, 0, time.UTC)) {
		startDate = time.Date(endDate.Year(), endDate.Month(), 1, 0, 0, 0, 0, time.UTC)
	}

	// Consultar los datos para el rango de fechas ajustado
	err := c.db.Where("meter_id IN (?) AND date BETWEEN ? AND ?",
		meterIDs, startDate, endDate).Find(&consumptions).Error
	if err != nil {
		return nil, err
	}
	fmt.Println("consumptions", consumptions)

	return consumptions, nil
}
