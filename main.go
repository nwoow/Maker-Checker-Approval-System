package main

import (
	"log"
	"maker_checker_sqlite/src/application"
	"maker_checker_sqlite/src/infrastructure/api"
	"maker_checker_sqlite/src/infrastructure/persistence"
	"net/http"
)

func main() {
	repo, err := persistence.NewSQLiteMessageRepository("messages.db")
	if err != nil {
		log.Fatalf("Failed to initialize SQLite repository: %v", err)
	}

	handler := &application.MessageHandler{MessageRepo: repo}
	controller := &api.MessageController{Handler: handler}

	http.HandleFunc("/send", controller.HandleSendMessage)
	http.HandleFunc("/approve", controller.HandleApproveMessage)

	log.Println("Starting server on :9090...")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
