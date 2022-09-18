package sippmadb

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type Database struct {
	Host       string `default:"localhost"`
	Port       string `default:"5432"`
	Username   string
	Password   string
	DBName     string
	Connection *sql.DB
}

func (db *Database) Connect() {
	var err error

	log.Print("Connecting to sippmadb")
	log.Printf("%v", db)
	pgsqlConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		db.Host, db.Port, db.Username, db.Password, db.DBName)
	db.Connection, err = sql.Open("postgres", pgsqlConnectionString)
	if err != nil {
		log.Fatalf("Could not connect to sippmadb: %s", err)
	}

	if err = db.Connection.Ping(); err != nil {
		panic(err)
	}
}

func (db *Database) ListUsers() {
	var firstname, lastname string
	var id int

	rows, err := db.Connection.Query("SELECT id, firstname, lastname FROM patients")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &firstname, &lastname)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, firstname, lastname)

	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func (db *Database) FindParents() {
	//TODO
}
