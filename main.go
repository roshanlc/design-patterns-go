package main

import "database/sql"

// Example of a basic singleton pattern
type Database struct {
	Connection *sql.DB
}

var instance *Database

func GetInstance() *Database {

	if instance != nil {
		return instance
	}

	// open an actual db connection and return it
	conn, err := sql.Open("connection string")

	if err != nil {
		panic("could not establish conn")
	}

	instance = &Database{Connection: conn}
	return instance
}

func main() {

	var dbConn = GetInstance()
	_ = dbConn

}
