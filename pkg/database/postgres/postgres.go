package postgres

import (
	"fmt"
	"go1/internal/config"
	"go1/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewPostgresDB(cfg config.PostgresConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host, cfg.UserName, cfg.Password, cfg.DbName, cfg.Port, cfg.SslMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&entity.Product{},
		&entity.Cart{},
		&entity.CartItem{},
		&entity.Order{},
		&entity.OrderItem{},
	)
	if err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}

	if !db.Migrator().HasConstraint(&entity.CartItem{}, "Cart") {
		if err := db.Migrator().CreateConstraint(&entity.CartItem{}, "Cart"); err != nil {
			return nil, fmt.Errorf("failed to create cart constraint: %w", err)
		}
	}

	return db, nil
}
