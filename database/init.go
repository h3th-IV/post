package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// Init dB
// Close dB
// Ping dB
// Create Tables
// Write Queries

var DB *sql.DB

// init the dB and try connecting
func InitDB() error {
	//constring for postgreSQL
	//host=your_host port=your_port user=your_user password=your_password dbname=your_dbname sslmode=disable
	dataSrcName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PG_HOST"), os.Getenv("PG_PORT"), os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"), os.Getenv("PG_DBNAME"))

	//log.Printf("Data Source Name: %s", dataSrcName) //dBugging
	var err error
	DB, err = sql.Open("postgres", dataSrcName)
	if err != nil {
		log.Fatal(err)
	}
	return DB.Ping()
}

func CloseDB() {
	DB.Close()
}

//secure socket not used since in Development env
