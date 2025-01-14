package backend

import (
	"encoding/json"
	"net/http"

	db "backend/db"
	models "backend/models"
	queue "backend/queue"

	"log"

	"github.com/gorilla/mux"
)

func ExecuteCode(w http.ResponseWriter, r *http.Request) {
	var req models.CodeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Save request in MongoDB
	id := db.SaveExecutionRequest(req)

	log.Printf("ID: %s \n", id)

	// Send task to RabbitMQ
	queue.PublishTask(req, id)

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"task_id": id})
}

func GetResult(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	result, err := db.GetExecutionResult(id)
	if err != nil {
		http.Error(w, "Result not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
