package main

import (
	"database/sql"
	"fmt"
	"os"
)

// Example of a basic prototype pattern
type DatabaseConfig struct {
	DSN string
}

type Database struct {
	Conn *sql.DB
}

func (dc *DatabaseConfig) Clone() Database {

	conn, err := sql.Open(dc.DSN, "")
	if err != nil {
		fmt.Println(err)
		return Database{}
	}

	return Database{Conn: conn}
}

func main() {

	dbConfig := DatabaseConfig{os.Getenv("DB_CONN_STRING")}

	db1 := dbConfig.Clone()
	db2 := dbConfig.Clone()

	_, _ = db1, db2

}
