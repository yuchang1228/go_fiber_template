package main

import (
	"app/internal/tasks"

	"github.com/RichardKnop/machinery/v2/log"
)

func main() {
	consumerTag := "machinery_worker"

	server, err := tasks.StartServer()
	if err != nil {
		log.FATAL.Fatalln("Could not start server:", err)
	}

	// The second argument is a consumer tag
	// Ideally, each worker should have a unique tag (worker1, worker2 etc)
	worker := server.NewWorker(consumerTag, 0)

	if err := worker.Launch(); err != nil {
		log.FATAL.Fatalln("Could not launch worker:", err)
	}
	log.INFO.Println("Worker is running...")
}
