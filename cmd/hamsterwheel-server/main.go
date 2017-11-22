package main

import (
	"log"

	"github.com/amerine/hamsterwheel/cmd/hamsterwheel-server/internal/database"
	"github.com/joeshaw/envdecode"
	_ "github.com/lib/pq"
)

func main() {
	var cfg config

	if err := envdecode.StrictDecode(&cfg); err != nil {
		log.Fatal(err)
	}

	_, err := database.New(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
}
