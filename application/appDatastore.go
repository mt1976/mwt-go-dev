package application

import (
	"database/sql"
	"fmt"
	"log"

	globals "github.com/mt1976/mwt-go-dev/globals"
)

// DataStoreConnect connects application to its datastore database
func DataStoreConnect() (*sql.DB, error) {
	// Connect to SQL Server DB
	mssqlConfig := GetProperties(globals.DATASTORECONFIG)

	server := mssqlConfig["server"]
	user := mssqlConfig["user"]
	password := mssqlConfig["password"]
	port := mssqlConfig["port"]
	database := mssqlConfig["database"]

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
		server, user, password, port, database)
	dbInstance, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	stmt, err := dbInstance.Prepare("select @@version")
	row := stmt.QueryRow()
	var result string

	err = row.Scan(&result)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	//fmt.Printf("%s\n", result)
	return dbInstance, err
}
