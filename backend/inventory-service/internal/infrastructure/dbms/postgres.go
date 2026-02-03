package dbms

import (
	"fmt"
	"log"
	"os"

	"inventory-service/internal/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	envFile := fmt.Sprintf(".env.%s", env)

	if err := godotenv.Load(envFile); err != nil {
		log.Printf("⚠️  Tidak dapat memuat file %s, mencoba .env default...", envFile)
		_ = godotenv.Load(".env")
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	sslmode := os.Getenv("SSL_MODE")
	tz := os.Getenv("TIMEZONE")

	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s TimeZone=%s",
		user, pass, host, port, name, sslmode, tz,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	if err := db.AutoMigrate(
		&model.Inventory{},
		&model.InventoryReservation{},
	); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	log.Println("✅ Database PostgreSQL connected successfully.")
	return db
}
