package main

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"

	"github.com/madasatya6/gin-gonic/config"
	"github.com/madasatya6/gin-gonic/internal/app"
)

/**
* GO CLEAN ARCHITECTURE
* 
* @author Mada Satya Bayu Ambika
* @version 2.12
* @link https://github.com/madasatya6
* @framework gin-gonic
* @development of Evrone
*
* @access public
* @note change "madasatya6/gin-gonic" according to the project name in all .go files
*/

func main() {
	// Configuration
	var cfg config.Config
	if err := cleanenv.ReadConfig("./config/config.yml", &cfg); err != nil {
		log.Fatalf("Config error: %s", err)
	}
	
	// Run
	app.Run(&cfg)
}
