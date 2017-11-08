package main

import (
	"log"

	"github.com/eriktate/simple/todo"
	"github.com/eriktate/simple/todo/roach"
)

func main() {
	host := "localhost"
	port := 26257
	user := "root"
	password := ""
	dbName := "todo"

	repo, err := roach.NewRepo(host, user, password, dbName, uint(port))
	if err != nil {
		log.Printf("Failed to get roach repo: %s", err)
	}

	if err := repo.RunQueryFile("../../roach/init.sql"); err != nil {
		log.Printf("Failed to run query file: %s", err)
	}

	t := todo.Todo{
		Title:       "Some title",
		Description: "Awesome description",
	}

	id, err := repo.Create(t)
	if err != nil {
		log.Printf("Failed to create: %s", err)
	}

	log.Printf("ID generated: %s", id)

	id2, err := repo.Create(t)
	if err != nil {
		log.Printf("Failed to create: %s", err)
	}

	todos, err := repo.GetAll()
	if err != nil {
		log.Printf("Failed to get todos: %s", err)
	}

	log.Printf("Todos: %+v", todos)

	deleted, err := repo.Delete(id2)
	if err != nil {
		log.Printf("Failed to delete todo: %s", err)
	}

	log.Printf("Deleted: %+v", deleted)

	todos, _ = repo.GetAll()

	log.Printf("Todos again: %+v", todos)
}
