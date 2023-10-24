package main

import (
	"apiassignment/config"
	"apiassignment/dal"
	"apiassignment/router"
	"log"
)

func main() {
	// Load the database configuration
	dbConfig, err := config.ConfigIniti("config/mysqldb.json")
	if err != nil {
		panic(err)
	}
	if err := dal.SqlDbInit(dbConfig); err != nil {
		log.Println(err)
	}

	router.RouterStart()

}
