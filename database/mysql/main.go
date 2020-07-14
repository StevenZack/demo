package main

import (
	"log"

	"github.com/StevenZack/demo/database/mysql/internal/storage"
)

func main() {
	table := storage.NewTable("user")
	e := table.Add(storage.NewRow([]string{
		"asd",
	}))
	if e != nil {
		log.Println(e)
		return
	}
}
