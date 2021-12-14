package main

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"

	"github.com/madasatya6/gin-gonic/config"
	"github.com/madasatya6/gin-gonic/internal/app"
)

func main() {
	// Configuration
	var cfg config.Config

	err := cleanenv.ReadConfig("./config/config.yml", &cfg)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(&cfg)
}
