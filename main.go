package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	LoadEnv()
	server := NewAPIServer(":3000")
	server.Run()
}

func getPort() string{
	port := os.Getenv("DB_PORT")
	if port == ""{
		return ":3000"
	}

	return ":" + port
}

func LoadEnv(){
	if err := godotenv.Load(); err != nil{
		log.Fatal("Error load .env file")
	}
}