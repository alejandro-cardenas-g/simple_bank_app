package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/alejandro-cardenas-g/simple_bank_app/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")

	if err != nil {
		log.Fatal("cannot load config")
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db")
	}
	testQueries = New(testDB)

	os.Exit(m.Run())
}
