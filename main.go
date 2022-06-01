package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Todo struct {
	Id      int    `json:"Id"`
	Title   string `json:"Title"`
	IsDone  bool   `json:"IsDone"`
	Content string `json:"Content"`
}

type Todos []Todo

var todos Todos

func getAllTodos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(todos)
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	var todo1 Todo

	err := json.NewDecoder(r.Body).Decode(&todo1)
	if err != nil {
		log.Fatal(err)
	}

	todos = append(todos, todo1)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	var todo1 Todo

	err := json.NewDecoder(r.Body).Decode(&todo1)
	if err != nil {
		log.Fatal(err)
	}

	for i, v := range todos {
		if v.Id == todo1.Id {
			todos = append(todos[:i], todos[i+1:]...)
		}
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", getAllTodos)
	myRouter.HandleFunc("/create", addTodo).Methods("POST")
	myRouter.HandleFunc("/delete", deleteTodo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
