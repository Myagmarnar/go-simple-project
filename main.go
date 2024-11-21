package main

import (
	"log"
	"net/http"
	"github.com/Myagmarnar/go-simple-project/handlers"
)

func main() {
    http.HandleFunc("/todos", handlers.GetTodos)
    http.HandleFunc("/todos/add", handlers.AddTodo)

    log.Println("Starting server on :8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}