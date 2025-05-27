package main

//javierstudent1
//d2I5XnbKcaY5b3rF
import (
	"backend/ecommerce-api/db"
	"backend/ecommerce-api/routes"
	"log"
	"net/http"
)

func main() {
	// Connect to the database
	log.Println("Connecting to db")
	db.ConnectMongo()

	// Init routes
	router := routes.SetupRoutes()

	// Start server
	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
