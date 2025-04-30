package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Initialize() {
	err := OpenDatabase()
	if err != nil {
		log.Printf("A new error %v", err)
	}
	defer CloseDatabase()
	fmt.Println("Done...")
}

func CreateTable() {
	create :=
		`CREATE TABLE IF NOT EXISTS users(
     userid bigint,
	 email varchar(100),
	 password varchar(100)
	);`

	_, err := DB.Exec(create)
	if err != nil {
		log.Printf("A new error %v", err)
	}

	insert :=
		`INSERT INTO users VALUES
	(1, 'user1@example.com', 'password1'),
	(2,'user2', 'user2@example.com', 'password2'),
	(3, 'user3@example.com', 'password3'),
	(4, 'user4@example.com', 'password4'),
	(5, 'user5@example.com', 'password5');`

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

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	CreateTable()
	return nil

}

func CloseDatabase() error {
	return DB.Close()
}

// func ValidateCredentials(username string, password string) error {
// 	query := "SELECT username, password FROM users WHERE username=$1"
// 	row := DB.QueryRow(query, username)

// 	var retrivedPassword, retrivedUsername string
// 	err := row.Scan(&retrivedUsername, &retrivedPassword)

// 	fmt.Printf("retrived username: %s", retrivedUsername)
// 	fmt.Printf("retrived password: %s", retrivedPassword)
// 	fmt.Printf(" username: %s", username)
// 	if err != nil {
// 		return errors.New("credentials invalid")
// 	}

// 	if retrivedUsername == username {
// 		if retrivedPassword == password {
// 			return nil
// 		}
// 	} else {
// 		return errors.New("Invalid UserName")
// 	}
// 	return nil

// }
