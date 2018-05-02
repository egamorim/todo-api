package main

import (
	"log"
	"net/http"

	"github/egamorim/todo-api/domain"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/todo", domain.GetTodoList).Methods("GET")
	router.HandleFunc("/todo/{id}", domain.GetTodo).Methods("GET")
	router.HandleFunc("/todo", domain.CreateTodo).Methods("POST")
	router.HandleFunc("/todo/{id}", domain.DeleteTodo).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
