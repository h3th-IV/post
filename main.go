package main

import (
	"log"

	"github.com/h3th-IV/postman/dataBase"
	"github.com/joho/godotenv"
)

func main() {
	//lod envroment varible
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	err = dataBase.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	defer dataBase.CloseDB()
	if err := dataBase.CreateTable(); err != nil {
		log.Fatal(err)
	}

	if err := dataBase.InputRecords("Invicible", "MJ"); err != nil {
		log.Fatal(err)
	}
	if err := dataBase.InputRecords("King's Disease", "Nas"); err != nil {
		log.Fatal(err)
	}
}
