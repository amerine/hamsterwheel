package main

import (
	"log"

	"github.com/amerine/hamsterwheel/cmd/hamsterwheel-server/internal/database"
	"github.com/joeshaw/envdecode"
	_ "github.com/lib/pq"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	bindata "github.com/mattes/migrate/source/go-bindata"
)

func main() {
	var cfg config

	if err := envdecode.StrictDecode(&cfg); err != nil {
		log.Fatal(err)
	}

	db, err := database.New(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	migrDBDriver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	migrAssets := bindata.Resource(database.AssetNames(),
		func(name string) ([]byte, error) {
			return database.Asset(name)
		})

	migrSrcDriver, err := bindata.WithInstance(migrAssets)
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithInstance("<go-bindata>", migrSrcDriver, "doesnt-matter", migrDBDriver)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Start: Running Migrations")
	err = m.Up()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Done: Running Migrations")
}
