package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db, _ = gorm.Open("mysql", "root:root@/todolist?charset=utf8&parseTime=True&loc=Local")

type TodoItemModel struct {
	Id          int `gorm:"primary_key"`
	Description string
	Completed   bool
}

func Home(w http.ResponseWriter, r *http.Request) {
	log.Info("API accessed")
	w.Header().Set("content-type", "application/json")
	_, err := io.WriteString(w, `{"msg": "Hello World"}`)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	description := r.FormValue("description")
	log.WithFields(log.Fields{"description": description}).Info("Add new TodoItem. Saving to database.")

	todo := &TodoItemModel{Description: description, Completed: false}
	db.Create(&todo)
	result := db.Last(&todo)

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(result.Value)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	// Get URL parameter from mux
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Test if the TodoItem exist in DB
	err := GetItemByID(id)
	if !err {
		w.Header().Set("Content-Type", "application/json")
		_, err := io.WriteString(w, `{"updated": false, "error": "Record Not Found"}`)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		completed, _ := strconv.ParseBool(r.FormValue("completed"))
		log.WithFields(log.Fields{"Id": id, "Completed": completed}).Info("Updating TodoItem")
		todo := &TodoItemModel{}
		db.First(&todo, id)
		todo.Completed = completed
		db.Save(&todo)

		w.Header().Set("Content-Type", "application/json")
		_, err := io.WriteString(w, `{"updated": true}`)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	// Get URL parameter from mux
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Test if the TodoItem exist in DB
	err := GetItemByID(id)
	if !err {
		w.Header().Set("Content-Type", "application/json")
		_, err2 := io.WriteString(w, `{"deleted": false, "error": "Record Not Found"}`)
		if err2 != nil {
			log.Fatal(err2)
		}
	} else {
		log.WithFields(log.Fields{"Id": id}).Info("Deleting TodoItem")
		todo := &TodoItemModel{}
		db.First(&todo, id)
		db.Delete(&todo)

		w.Header().Set("Content-Type", "application/json")
		_, err2 := io.WriteString(w, `{"deleted": true}`)
		if err2 != nil {
			log.Fatal(err2)
		}
	}
}

func GetItemByID(Id int) bool {
	todo := &TodoItemModel{}
	result := db.First(&todo, Id)

	if result.Error != nil {
		log.Warn("TodoItem not found in database")

		return false
	}

	return true
}

func GetCompletedItems(w http.ResponseWriter, r *http.Request) {
	log.Info("Get completed TodoItems")

	completedTodoItems := GetTodoItems(true)
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(completedTodoItems)
	if err != nil {
		log.Fatal(err)
	}
}

func GetIncompleteItems(w http.ResponseWriter, r *http.Request) {
	log.Info("Get Incomplete TodoItems")

	IncompleteTodoItems := GetTodoItems(false)
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(IncompleteTodoItems)
	if err != nil {
		log.Fatal(err)
	}
}

func GetTodoItems(completed bool) interface{} {
	var todos []TodoItemModel

	TodoItems := db.Where("completed = ?", completed).Find(&todos).Value

	return TodoItems
}

func main() {
	defer db.Close()

	db.Debug().DropTableIfExists(&TodoItemModel{})
	db.Debug().AutoMigrate(&TodoItemModel{})

	log.Info("Starting server at PORT: 5000")

	router := mux.NewRouter()

	router.HandleFunc("/home", Home).Methods("GET")
	router.HandleFunc("/todo-completed", GetCompletedItems).Methods("GET")
	router.HandleFunc("/todo-incomplete", GetIncompleteItems).Methods("GET")
	router.HandleFunc("/todo", CreateItem).Methods("POST")
	router.HandleFunc("/todo/{id}", UpdateItem).Methods("POST")
	router.HandleFunc("/todo/{id}", DeleteItem).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))
}
