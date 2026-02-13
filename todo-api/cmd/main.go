package main

import (
	"github.com/pemtshewang/todo-api/internal/task"
	"log"
	"net/http"
)

func main() {
	repo := task.NewRepository()
	service := task.NewService(repo)
	handler := task.NewHandler(service)

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.CreateTask(w, r)
		case http.MethodGet:
			handler.GetTasks(w, r)
		case http.MethodDelete:
			handler.DeleteTask(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server x running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
