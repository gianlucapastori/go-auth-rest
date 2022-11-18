package main

import (
	"fmt"
	"log"

	"github.com/gianlucapastori/nausicaa/cfg"
)

func main() {
	v, err := cfg.New("config")
	if err != nil {
		log.Panicf("could not create a new viper instance: %v\n", err)
	}

	cfg, err := cfg.Parse(v)
	if err != nil {
		log.Panicf("could not create parse to config struct: %v\n", err)
	}

	fmt.Print(cfg)
}
