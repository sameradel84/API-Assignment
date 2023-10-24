package dal

import (
	"apiassignment/config"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var dbs *sql.DB

func SqlDbInit(cnf *config.Config) error {

	// DNS := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True", cnf.User, cnf.Password, cnf.Server, cnf.Db)

	// fmt.Println(DNS)

	var err error

	dbs, err = sql.Open("mysql", "root:samer@tcp(127.0.0.1:3306)/petsDB")

	if err != nil {
		log.Println("Error in opening SQL Connection", err.Error())
		return err
	}
	log.Println("SQL connection successful")
	return nil
}

func GetSQLDBConn() *sql.DB {
	return dbs
}

func CloseSQLConn() {
	dbs.Close()
}
