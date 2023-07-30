package repositories

import (
	"fmt"
	"github.com/cantillo16/bia_energy/src/models"
	"github.com/cantillo16/bia_energy/src/utils"
	"gorm.io/gorm"
	"sort"
	"time"
)

type ConsumptionRepository interface {
	GetConsumptionWeekly(meterIDs []string, startDate, endDate time.Time) ([]models.Consumption, error)
	GetMonthlyConsumption(meterIDs []string, startDate, endDate time.Time) ([]models.Consumption, error)
	GetConsumptionDaily(meterIDs []string, startDate, endDate time.Time) ([]models.Consumption, error)
	GetLastConsumption() ([]models.Consumption, error)
}

type ConsumptionRepositoryImpl struct {
	db *gorm.DB
}

func NewConsumptionRepository(db *gorm.DB) *ConsumptionRepositoryImpl {
	return &ConsumptionRepositoryImpl{db}

}

func (c *ConsumptionRepositoryImpl) GetMonthlyConsumption(
	meterIDs []string, startDate, endDate time.Time) ([]models.Consumption, error) {
	var consumptions []models.Consumption

	// Convertir las fechas a la zona horaria UTC
	startDate, endDate = utils.SetStartAndEndOfMonth(startDate.UTC(), endDate.UTC())

	// Convertir las fechas a cadenas con formato de fecha y hora en UTC
	startDateStr := startDate.UTC().Format("2006-01-02 15:04:05")
	endDateStr := endDate.UTC().Format("2006-01-02 15:04:05")

	err := c.db.Where("meter_id IN (?) AND date BETWEEN ? AND ?",
		meterIDs, startDateStr, endDateStr).Find(&consumptions).Error
	if err != nil {
		return nil, err
	}

	// Ordenar las consumptions por fecha antes de devolverlas
	sort.SliceStable(consumptions, func(i, j int) bool {
		return consumptions[i].Date.Before(consumptions[j].Date)
	})

	return consumptions, nil
}

func (c *ConsumptionRepositoryImpl) GetConsumptionWeekly(
	meterIDs []string, startDate, endDate time.Time) ([]models.Consumption, error) {
	var consumptions []models.Consumption

	startDate, endDate = utils.SetStartAndEndWeek(startDate, endDate)

	// Convertir las fechas a cadenas con formato de fecha y hora en UTC
	startDateStr := startDate.UTC().Format("2006-01-02 15:04:05")
	endDateStr := endDate.UTC().Format("2006-01-02 15:04:05")

	fmt.Println("startDateStr: ", startDateStr)
	fmt.Println("endDateStr: ", endDateStr)

	err := c.db.Where("meter_id IN (?) AND date BETWEEN ? AND ?",
		meterIDs, startDateStr, endDateStr).Find(&consumptions).Error
	if err != nil {
		return nil, err
	}
	return consumptions, nil
}

func (c *ConsumptionRepositoryImpl) GetConsumptionDaily(
	meterIDs []string, startDate, endDate time.Time) ([]models.Consumption, error) {
	var consumptions []models.Consumption

	startDate, endDate = utils.SetStartAndEndDay(startDate.UTC(), endDate.UTC())

	// Convertir las fechas a cadenas con formato de fecha y hora en UTC
	startDateStr := startDate.UTC().Format("2006-01-02 15:04:05")
	endDateStr := endDate.UTC().Format("2006-01-02 15:04:05")

	fmt.Println("startDateStr: ", startDateStr)
	fmt.Println("endDateStr: ", endDateStr)

	err := c.db.Where("meter_id IN (?) AND date BETWEEN ? AND ?",
		meterIDs, startDateStr, endDateStr).Find(&consumptions).Error
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
