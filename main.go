package main

import (
	"log"

	"github.com/h3th-IV/postman/database"
	"github.com/joho/godotenv"
)

func main() {
	//lod envroment varible
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	err = database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	defer database.CloseDB()
	//perfom CRUD
	err = database.InputRecords("Heavy", "Drake")
	if err != nil {
		log.Println(err)
	}

	// dataBase.UpdateRecords("New Album", "Nas")
}
