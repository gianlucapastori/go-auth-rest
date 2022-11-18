package main

import (
	"log"

	"github.com/gianlucapastori/nausicaa/cfg"
	"github.com/gianlucapastori/nausicaa/internal/server"
	"github.com/gianlucapastori/nausicaa/pkg/database"
	"github.com/gianlucapastori/nausicaa/pkg/logger"
)

func main() {
	v, err := cfg.New("config")
	if err != nil {
		log.Panic(err)
	}

	cfg, err := cfg.Parse(v)
	if err != nil {
		log.Panic(err)
	}

	db, err := database.Connect(cfg)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	sugar, err := logger.New()
	if err != nil {
		log.Panic(err)
	}

	var s *server.Server

	if err := s.New(db, cfg, sugar).Serve(); err != nil {
		log.Panic(err)
	}
}
