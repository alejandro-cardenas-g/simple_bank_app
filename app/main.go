package main

import (
	"database/sql"
	"log"

	"github.com/alejandro-cardenas-g/simple_bank_app/api"
	db "github.com/alejandro-cardenas-g/simple_bank_app/db/sqlc"
	"github.com/alejandro-cardenas-g/simple_bank_app/util"
	_ "github.com/lib/pq"
)

func main() {

	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db")
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal("Cannot start server")
	}

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("Cannot start server")
	}
}
