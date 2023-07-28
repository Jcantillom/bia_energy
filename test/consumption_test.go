package test

import (
	"github.com/cantillo16/bia_energy/src/models"
	"github.com/cantillo16/bia_energy/src/services"
	"testing"
	"time"
)

func TestGetConsumptionMonthly(t *testing.T) {
	// Crea una instancia del mock del repositorio
	repoMock := &MockConsumptionRepository{
		GetConsumptionMonthlyFunc: func(meterIDs []int, startDate, endDate time.Time) ([]models.Consumption, error) {
			return []models.Consumption{
				{ID: "1", MeterID: 1, ActiveEnergy: 100, ReactiveEnergy: 50, Date: time.Now()},
				{ID: "2", MeterID: 1, ActiveEnergy: 200, ReactiveEnergy: 100, Date: time.Now()},
			}, nil
		},
	}

	// Crea una instancia del servicio y le pasa el mock del repositorio
	service := services.NewConsumptionService(repoMock)

	// Llama al método del servicio que se quiere probar
	consumptions, err := service.GetConsumptionMonthly([]int{1}, time.Now(), time.Now())

	// Verifica que no haya error
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	// Verifica que el resultado sea el esperado
	if len(consumptions) != 2 {
		t.Errorf("Expected 2, got %d", len(consumptions))
	}
}

func TestGetConsumptionDaily(t *testing.T) {
	// Crea una instancia del mock del repositorio
	repoMock := &MockConsumptionRepository{
		GetConsumptionDailyFunc: func(meterIDs []int, startDate, endDate time.Time) ([]models.Consumption, error) {
			return []models.Consumption{
				{ID: "1", MeterID: 1, ActiveEnergy: 100, ReactiveEnergy: 50, Date: time.Now()},
				{ID: "2", MeterID: 1, ActiveEnergy: 200, ReactiveEnergy: 100, Date: time.Now()},
			}, nil
		},
	}

	// Crea una instancia del servicio y le pasa el mock del repositorio
	service := services.NewConsumptionService(repoMock)

	// Llama al método del servicio que se quiere probar
	consumptions, err := service.GetConsumptionDaily([]int{1}, time.Now(), time.Now())

	// Verifica que no haya error
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	// Verifica que el resultado sea el esperado
	if len(consumptions) != 2 {
		t.Errorf("Expected 2, got %d", len(consumptions))
	}
}

func TestGetConsumptionWeekly(t *testing.T) {
	// Crea una instancia del mock del repositorio
	repoMock := &MockConsumptionRepository{
		GetConsumptionWeeklyFunc: func(meterIDs []int, startDate, endDate time.Time) ([]models.Consumption, error) {
			return []models.Consumption{
				{ID: "1", MeterID: 1, ActiveEnergy: 100, ReactiveEnergy: 50, Date: time.Now()},
				{ID: "2", MeterID: 1, ActiveEnergy: 200, ReactiveEnergy: 100, Date: time.Now()},
			}, nil
		},
	}

	// Crea una instancia del servicio y le pasa el mock del repositorio
	service := services.NewConsumptionService(repoMock)

	// Llama al método del servicio que se quiere probar
	consumptions, err := service.GetConsumptionWeekly([]int{1}, time.Now(), time.Now())

	// Verifica que no haya error
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	// Verifica que el resultado sea el esperado
	if len(consumptions) != 2 {
		t.Errorf("Expected 2, got %d", len(consumptions))
	}
}
