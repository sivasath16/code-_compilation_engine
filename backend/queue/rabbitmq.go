package backend

import (
	"encoding/json"
	"log"

	db "backend/db"
	executor "backend/executor"
	models "backend/models"

	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var ch *amqp.Channel
var queue amqp.Queue

// ConnectRabbitMQ establishes a connection to RabbitMQ
func ConnectRabbitMQ() {
	var err error
	conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	ch, err = conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	queue, err = ch.QueueDeclare(
		"code_tasks",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	log.Println("RabbitMQ connection established.")
}

// CloseRabbitMQ closes the connection and channel
func CloseRabbitMQ() {
	if ch != nil {
		ch.Close()
	}
	if conn != nil {
		conn.Close()
	}
}

func StartTaskConsumer() {
	msgs, err := ch.Consume(
		queue.Name, // Queue name
		"",         // Consumer
		true,       // Auto-ack
		false,      // Exclusive
		false,      // No-local
		false,      // No-wait
		nil,        // Args
	)
	if err != nil {
		log.Fatalf("Failed to register consumer: %v", err)
	}

	// Consume tasks
	log.Println("Task consumer started...")
	for msg := range msgs {
		var task models.Task
		err := json.Unmarshal(msg.Body, &task)
		if err != nil {
			log.Printf("Failed to parse task: %v", err)
			continue
		}

		log.Printf("MSG: %v \n", msg.Body)

		log.Printf("Processing task ID: %s", task.ID)

		// Execute the code in a Docker container
		output, err := executor.ExecuteTask(task.Code, task.Language)
		result := models.ExecutionResult{
			ID:     task.ID,
			Output: output,
		}
		if err != nil {
			log.Printf("Task execution failed: %v", err)
			result.Error = err.Error()
		}

		// Save the result in MongoDB
		db.SaveExecutionResult(result)
		log.Printf("Saving result for Task ID: %s, Output: %s", result.ID, result.Output)
		log.Printf("Task ID %s completed", task.ID)
	}
}

// PublishTask sends a task to the RabbitMQ queue
func PublishTask(req models.CodeRequest, taskID string) {
	task := models.Task{
		ID:       taskID,
		Code:     req.Code,
		Language: req.Language,
	}
	body, err := json.Marshal(task)
	if err != nil {
		log.Printf("Failed to serialize task: %v", err)
		return
	}

	err = ch.Publish(
		"",         // Exchange
		queue.Name, // Routing key
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Printf("Failed to publish task: %v", err)
	}
}
