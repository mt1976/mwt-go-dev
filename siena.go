package main

import (
	"database/sql"
	"fmt"
	"log"
)

func sienaConnect() (*sql.DB, error) {
	// Connect to SQL Server DB
	mssqlConfig := getProperties(cSQL_CONFIG)

	server := mssqlConfig["server"]
	user := mssqlConfig["user"]
	password := mssqlConfig["password"]
	port := mssqlConfig["port"]
	database := mssqlConfig["database"]
	/*
		fmt.Println("SQL SERVER - CONNECTING...")
		fmt.Println("Server     :", server)
		fmt.Println("User       :", user)
		fmt.Println("Password   :", strings.Repeat("*", len(password)))
		fmt.Println("Port       :", port)
		fmt.Println("Database   :", database)
		fmt.Println("")
	*/
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
		server, user, password, port, database)
	dbInstance, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	//fmt.Printf("Connected!\n")
	//fmt.Println("")

	//defer dbInstance.Close()

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
