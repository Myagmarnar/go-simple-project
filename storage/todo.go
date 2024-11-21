package storage

import (
	"github.com/Myagmarnar/go-simple-project/models"
)

var todos []models.Todo
var nextID = 1

func GetTodos() []models.Todo {
	return todos
}

func AddTodo(todo models.Todo) {
	todo.ID = nextID
	nextID++
	todos = append(todos, todo)
}
