package connection

import (
	"fmt"
	"github.com/cantillo16/bia_energy/src/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func Connect() *gorm.DB {
	_ = godotenv.Load()
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	return db
}

func Migrate() {
	db := Connect()
	db.AutoMigrate(&models.Consumption{})
	fmt.Println("Database migrated successfully! ▶️")
}
