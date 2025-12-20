package main

import (
	"DSCMailer/internal/server"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading the .env file")
		return
	}
}

func main() {

	router := server.NewRouter()

	fmt.Println("Server listening on http://127.0.0.1:3000")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
