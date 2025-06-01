package main

import (
	"log"

	"github.com/faizinahsan/academic-system/config"
	"github.com/faizinahsan/academic-system/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
