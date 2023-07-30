package importer

import (
	"encoding/csv"
	"fmt"
	"github.com/cantillo16/bia_energy/src/models"
	"gorm.io/gorm"
	"os"
	"strconv"
	"time"
)

func ImportCSVToDB(db *gorm.DB, filePath string) error {
	fmt.Println("Cargando CSV a la base de datos, esto tomara 3 min y 22 Seg, " +
		"espere el mensaje de coonfirmacion ...⏳ ")

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// create a new reader
	reader := csv.NewReader(file)

	_, _ = reader.Read() // skip header

	// read all records and load into db
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		// parse record
		id := record[0]
		meterID := record[1]
		activeEnergy, _ := strconv.ParseFloat(record[2], 64)
		reactiveEnergy, _ := strconv.ParseFloat(record[3], 64)
		capacitiveReactive, _ := strconv.ParseFloat(record[4], 64)
		solar, _ := strconv.ParseFloat(record[5], 64)

		date, err := parseDate(record[6])
		if err != nil {
			fmt.Println(err)
			continue
		}

		// create consumption
		consumption := models.Consumption{
			ID:                 id,
			MeterID:            meterID,
			ActiveEnergy:       activeEnergy,
			ReactiveEnergy:     reactiveEnergy,
			CapacitiveReactive: capacitiveReactive,
			Solar:              solar,
			Date:               date,
			CreateAt:           time.Now(),
			UpdateAt:           time.Now(),
		}

		// insert consumption
		err = db.Create(&consumption).Error
		if err != nil {
			return err
		}
	}
	fmt.Println("CSV cargado exitosamente a la base de datos ✅")
	return nil
}

func parseDate(dateStr string) (time.Time, error) {
	if dateStr == "date" {
		return time.Time{}, nil
	}

	layout := "2006-01-02 15:04:05-07"
	parsedDate, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, nil
	}

	return parsedDate, nil
}
