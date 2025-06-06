package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type (
	Config struct {
		Environment string
		Postgres    PostgresConfig
		Http        HttpConfig
	}
	PostgresConfig struct {
		UserName string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		DbName   string `mapstructure:"dbname"`
		SslMode  string `mapstructure:"sslmode"`
	}

	HttpConfig struct {
		Host               string        `mapstructure:"host"`
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
	}
)

func Init() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env: %v", err) // Добавьте лог
		return nil, fmt.Errorf("error loading .env: %w", err)
	}
	var cfg Config
	setFromEnv(&cfg)

	return &cfg, nil
}

func setFromEnv(cfg *Config) {
	cfg.Postgres.UserName = os.Getenv("DB_USER")
	cfg.Postgres.Password = os.Getenv("DB_PASSWORD")
	cfg.Postgres.Host = os.Getenv("DB_HOST")
	cfg.Postgres.Port = os.Getenv("DB_PORT")
	cfg.Postgres.DbName = os.Getenv("DB_NAME")
	cfg.Postgres.SslMode = os.Getenv("DB_SSLMODE")

	cfg.Http.Port = os.Getenv("SERVER_PORT")
	cfg.Environment = os.Getenv("APP_ENV")
}
