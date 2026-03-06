package config

import (
	"database/sql"

	 _ "github.com/go-sql-driver/mysql"
)

func DBconnection() (*sql.DB, error){
	dbDriver := "mysql"
	dbUser := "root"
	dbpass := ""
	dbName := "golang_dendra"

	db, err := sql.Open(dbDriver, dbUser+":"+dbpass+"@/"+dbName)
	return db, err
}