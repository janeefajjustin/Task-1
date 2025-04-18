package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func CreateTable() {
	create :=
		`CREATE TABLE IF NOT EXISTS users(
     username varchar(100),
	 email varchar(100),
	 password varchar(100)
	):`
	
	_, err := DB.Exec(create)
	if err != nil {
		log.Printf("A new error %v", err)
	}

		insert :=
			`INSERT INTO users VALUES
	('user1', 'user1@example.com', 'password1'),
	('user2', 'user2@example.com', 'password2'),
	('user3', 'user3@example.com', 'password3'),
	('user4', 'user4@example.com', 'password4'),
	('user5', 'user5@example.com', 'password5');`

		_, err = DB.Exec(insert)
		if err != nil {
			log.Printf("A new error %v", err)
		}

}

func OpenDatabase() error {
	var err error
	DB, err = sql.Open("postgres", "user=postgres password=8976 dbname=TaskOne sslmode=disable")
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	CreateTable()
	return nil

}

func CloseDatabase() error {
	return DB.Close()
}
