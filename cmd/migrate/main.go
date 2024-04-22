package main

import (
	"log"
	"os"

	"github.com/celsopires1999/ecom/configs"
	"github.com/celsopires1999/ecom/internal/db"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	cfg := db.Config{
		DBHost:     configs.Envs.DBHost,
		DBPort:     configs.Envs.DBAPort,
		DBUser:     configs.Envs.DBUser,
		DBPassword: configs.Envs.DBPassword,
		DBName:     configs.Envs.DBName,
	}

	db, err := db.NewPostgresStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	v, d, _ := m.Version()
	log.Printf("Version: %d, dirty: %v", v, d)

	cmd := os.Args[len(os.Args)-1]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}

}
