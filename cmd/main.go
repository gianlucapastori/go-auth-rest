package main

import (
	"log"

	"github.com/gianlucapastori/nausicaa/cfg"
	"github.com/gianlucapastori/nausicaa/pkg/database"
)

func main() {
	v, err := cfg.New("config")
	if err != nil {
		log.Panicf("error while creating viper struct: %v\n", err)
	}

	cfg, err := cfg.Parse(v)
	if err != nil {
		log.Panicf("error while parsing config into struct: %v\n", err)
	}

	db, err := database.Connect(cfg)
	if err != nil {
		log.Panicf("error while connecting to database: %v\n", err)
	}
	defer db.Close()
}
