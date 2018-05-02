package domain

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var todoList []Todo

//GetTodoList return the todo list
func GetTodoList(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(all())
}

//GetTodo - get todo by id
func GetTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	todo := findById(params["id"])

	json.NewEncoder(w).Encode(&todo)
}

//CreateTodo create new
func CreateTodo(w http.ResponseWriter, r *http.Request) {

	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)

	saveTodo(todo)
	json.NewEncoder(w).Encode(all())
}

//DeleteTodo - remove by id
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range todoList {
		if item.ID == params["id"] {
			todoList = append(todoList[:index], todoList[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(todoList)
	}
}
