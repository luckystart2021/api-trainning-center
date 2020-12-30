package main

import (
	"api-trainning-center/database"
	"api-trainning-center/handlers"
	"log"
	"net/http"
)

func main() {
	db, err := database.Initialize()
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}
	defer db.Close()

	httpHandler := handlers.NewHandler(db)
	log.Println("Server started on: http://localhost:8080")
	http.ListenAndServe(":8080", httpHandler)
}
