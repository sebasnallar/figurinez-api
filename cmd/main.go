package main

import (
	"figurinez-api/internal/infrastructure"
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	infrastructure.InitializeConfig()

	infrastructure.InitializeLogger()

	infrastructure.Logger.Info("Logger initialized")

	db, err := infrastructure.InitializeDatabase()
	if err != nil {
		infrastructure.Logger.Fatal("Failed to initialize database", zap.Error(err))
	}
	deps := infrastructure.SetupDependencies(db)

	router := infrastructure.SetupRouter(deps)

	port := viper.GetInt("server.port")
	if err := router.Run(fmt.Sprintf(":%d", port)); err != nil {
		infrastructure.Logger.Fatal("Failed to run server", zap.Error(err))
	}
}
