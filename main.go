package main

import (
	"fmt"
	"github.com/shubham-prasad/awesome-wallet/util"
)

const (
	dbDriver = "postgres"
	configPath = "app.env"
)

func main() {
	//dbConn, err := sql.Open(dbDriver, dbSource)
	//if err != nil {
	//	log.Fatal("DB connection error > ", err)
	//}
	//storage := db.NewStorage(dbConn)
	config, err := util.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("config", config)
	fmt.Println("conn string", config.GetDbConnectionUrl())
}
