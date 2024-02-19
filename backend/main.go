package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"infy/db"
	"infy/routes"
	"infy/utils"
	"log"
)

func main() {
	fmt.Println("Loading .env file...")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Connecting to database...")
	db.InitMongo()
	defer db.CloseMongo()

	fmt.Println("Starting server...")
	port := ":" + utils.GetEnv("PORT", "8000")
	r := routes.InitRoutes()

	err = r.Run(port)
	if err != nil {
		log.Fatal(err)
	}
}
