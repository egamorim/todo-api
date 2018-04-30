package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var todoList []Todo

// Todo struct
type Todo struct {
	ID         string    `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty`
}

//GetTodoList return the todo list
func GetTodoList(w http.ResponseWriter, r *http.Request) {
	todoList = append(todoList, Todo{ID: "1", Name: "Task 1", CreateDate: time.Now()})
	todoList = append(todoList, Todo{ID: "2", Name: "Task 2", CreateDate: time.Now()})
	json.NewEncoder(w).Encode(todoList)
}

//GetTodo - get todo by id
func GetTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range todoList {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Todo{})
}

//CreateTodo create new
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = params["id"]
	todoList = append(todoList, todo)
	json.NewEncoder(w).Encode(todoList)
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/todo", GetTodoList).Methods("GET")
	router.HandleFunc("/todo/{id}", GetTodo).Methods("GET")
	router.HandleFunc("/todo/{id}", CreateTodo).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
