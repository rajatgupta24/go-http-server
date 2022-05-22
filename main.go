package main

import (
	"fmt"
	"net/http"
)

type ToDoItem struct {
	id     int
	todo   string
	isDone bool
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "Hello World")
}

func addElement(todos []ToDoItem, todo ToDoItem) ToDoItem {
	fmt.Println(todos, todo)
}

func main() {
	todos := []ToDoItem{
		{
			1,
			"Need to test my code",
			false,
		},
		{
			2,
			"Dockerise the app",
			false,
		},
	}

	// fileServer := http.FileServer(http.Dir("./static"))
	// http.Handle("/", fileServer)
	// http.HandleFunc("/hello", helloHandler)

	// fmt.Println("Server listening on PORT: 8080")
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	log.Fatal(err)
	// }

	todo := ToDoItem{3, "Deploy it on k3s", false}

	addElement(todos, todo)

	fmt.Println(todos)
}
