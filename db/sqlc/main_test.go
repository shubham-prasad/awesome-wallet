package db

import (
	"database/sql"
	"github.com/shubham-prasad/awesome-wallet/util"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var bankDb *Storage

const (
	dbDriver = "postgres"
	configPath = "../../app.env"
)

func TestMain(m *testing.M) {
	config, err := util.LoadConfig(configPath)
	if err != nil {
		log.Fatal("could not load config", err)
	}
	dbConn, err := sql.Open(dbDriver, config.GetDbConnectionUrl())
	if err != nil {
		log.Fatal("DB connection error > ", err)
	}
	testQueries = New(dbConn)
	bankDb = NewStorage(dbConn)
	os.Exit(m.Run())
}
