package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/Myagmarnar/go-simple-project/models"
    "github.com/Myagmarnar/go-simple-project/storage"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
    todos := storage.GetTodos()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todos)
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
    var todo models.Todo
    if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }
    storage.AddTodo(todo)
    w.WriteHeader(http.StatusCreated)
}
