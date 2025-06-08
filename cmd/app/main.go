package main

import (
	"github.com/faizinahsan/academic-system/config"
	"github.com/faizinahsan/academic-system/internal/app"
	"log"
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
