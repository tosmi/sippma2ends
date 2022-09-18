package main

import (
	"fmt"
	"github.com/tosmi/sippma2ends/pkg/sippmadb"
	"log"
)

func main() {
	fmt.Println("hello world")
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var database = sippmadb.Database{
		Host:     "localhost",
		Port:     "5432",
		Username: "sippma",
		Password: "sippma",
		DBName:   "sippma",
	}

	database.Connect()
	database.ListUsers()
}
