package main

import (
	"context"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/configs"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/registry"
	"github.com/khanghld27/kelvin-kart-challenge-api/cmd/services/common"
	"github.com/khanghld27/kelvin-kart-challenge-api/pkg/logger"
	"github.com/sirupsen/logrus"
	"os"
)

func loadConfig() configs.Config {
	var cfg configs.Config
	if err := envconfig.Process("", &cfg); err != nil {
		logrus.Fatal(err)
	}

	configs.SetConfig(&cfg)
	return cfg
}

func main() {
	ctx := context.TODO()
	cfg := loadConfig()
	logger.Init(cfg.LogLevel == string(logger.DebugLevel))
	common.InitDBConnection(cfg.PostgreSQL) // assumes you have this

	txnMiddleware := registry.TransactionMiddleware() // assumes you have this

	ctx = txnMiddleware.StartToolRequest(ctx)

	apiBaseURL := os.Getenv("EXTERNAL_API_URL")
	if apiBaseURL == "" {
		panic("EXTERNAL_API_URL not set")
	}

	importer := registry.InitializeProductImporter(apiBaseURL)

	if err := importer.ImportProducts(ctx); err != nil {
		fmt.Printf("Import failed: %v\n", err)
		os.Exit(1)
	}

	txnMiddleware.EndToolRequest(ctx)

	fmt.Println("Products imported successfully!")
}
