package main

/*
	Entry Point for App
*/

import (
	"log"

	"github.com/joho/godotenv"
)

func main(){
	//Load vars from .env into the application environment
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error Loading .env file")
	}
}
