package main

import (
	"fmt"
	"gemini-cli-go/cmd"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	key := os.Getenv("GEMINI_API_KEY")
	// fmt.Println("Hello, World!")
	// fmt.Println(key)

	if key == "" {
		fmt.Println("GEMINI_API_KEY is not set in environment variables")
		return
	}

	// Executing root command (cmd/root.go)
	cmd.Execute()
}
