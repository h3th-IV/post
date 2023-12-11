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
	//perfom CRUD
}
