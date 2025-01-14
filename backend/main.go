package main

import (
	"log"
	"net/http"

	d "backend/db"
	h "backend/handlers"
	queue "backend/queue"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Connect to MongoDB
	d.ConnectMongoDB()

	// Initialize RabbitMQ
	queue.ConnectRabbitMQ()
	defer queue.CloseRabbitMQ()

	// Start task consumer
	go queue.StartTaskConsumer()

	// Set up router
	r := mux.NewRouter()
	log.Printf("Calling ExecuteCode")
	r.HandleFunc("/api/execute", h.ExecuteCode).Methods("POST")
	r.HandleFunc("/api/results/{id}", h.GetResult).Methods("GET")

	// CORS configuration
	corsOptions := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	corsHeaders := handlers.AllowedHeaders([]string{"Content-Type"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	// Start server with CORS middleware
	log.Println("Backend running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(corsOptions, corsHeaders, corsMethods)(r)))
}
