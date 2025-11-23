package config

import (
	"rest/internal/utils/env"
	"time"
)

type config struct {
	Addr       string
	DbConfig   dbConfig
	AuthConfig authConfig
}

type dbConfig struct {
	Addr         string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  time.Duration
}

func Load() *config {
	return &config{
		Addr: env.GetString("ADDR", ":8080"),
		DbConfig: dbConfig{
			Addr:         env.GetString("DB_ADDR", "postgres://user:password@localhost:5433/ecommerce_dev?sslmode=disable"),
			MaxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			MaxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			MaxIdleTime:  env.GetDuration("DB_MAX_IDLE_TIME", time.Minute*15),
		},
	}
}

type authConfig struct {
	JWTSecret string
	JWTExpire time.Duration
}

var AuthConfig authConfig

func init() {
	AuthConfig = authConfig{
		JWTSecret: env.GetString("JWT_SECRET", ""),
		JWTExpire: env.GetDuration("JWT_EXPIRED_AT", time.Hour*24*7),
	}
}
