package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func main() {
	db, err := sql.Open("mysql", "root:root@/mysql")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	db.Exec("create database if not exists todolist;")
	db.Exec("use todolist;")
	db.Exec("create table todos(id int, todo varchar(255));")

	todos := []struct {
		id   int
		todo string
	}{
		{1, "Create a go-app"},
		{2, "Dockerize the app"},
		{3, "Deploy the app on k3s"},
	}

	stmt, err := db.Prepare("INSERT INTO todos(id, todo) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	for _, project := range todos {
		if _, err := stmt.Exec(project.id, project.todo); err != nil {
			log.Fatal(err)
		}
	}

	defer stmt.Close()

	rows, err := db.Query("select * from todos")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var (
			id   int64
			name string
		)
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		log.Printf("id %d name is %s\n", id, name)
	}

	defer rows.Close()

	http.HandleFunc("/home", Home)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
