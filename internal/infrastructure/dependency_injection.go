package infrastructure

import (
	"figurinez-api/internal/health"

	"gorm.io/gorm"
)

type Services struct {
}

type Handlers struct {
	HealthHandler *health.HealthHandler
}

type Dependencies struct {
	DB       *gorm.DB
	Services *Services
	Handlers *Handlers
}

func SetupDependencies(db *gorm.DB) *Dependencies {
	services := &Services{}
	handlers := &Handlers{
		HealthHandler: &health.HealthHandler{},
	}
	return &Dependencies{
		DB:       db,
		Services: services,
		Handlers: handlers,
	}
}
