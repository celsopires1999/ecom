package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/celsopires1999/ecom/cmd/api"
	"github.com/celsopires1999/ecom/configs"
	"github.com/celsopires1999/ecom/internal/db"
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

	initStorage(db)

	server := api.NewAPIServer(fmt.Sprintf(":%s", configs.Envs.Port), nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}
