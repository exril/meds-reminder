package api

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error While Loading .env File , Please Check if the file exists")
	}
	fmt.Println("Succesfully Loaded ENV Variables")
}