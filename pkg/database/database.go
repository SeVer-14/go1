package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Connect() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	return nil
}

func Migrate() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	dialect := os.Getenv("MIGRATIONS_DIALECT")
	if err := goose.SetDialect(dialect); err != nil {
		return err
	}

	dir := os.Getenv("MIGRATIONS_DIR")
	return goose.Up(sqlDB, dir)
}
