package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/janeefajjustin/task-1/models"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Initialize() {
	err := OpenDatabase()
	fmt.Println("DB Initializing....")
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
	(2, 'user2@example.com', 'password2'),
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

//Code from repo

func ValidateCredentials(u *models.User) error {
	query := "SELECT userid, password FROM users WHERE email=$1"


	//test
	fmt.Printf("u.Email %v",u.Email)
	fmt.Printf("u.Password %v",u.Password)
	fmt.Printf("u.ID %vb\n",u.ID)

	//add db
	row := DB.QueryRow(query, u.Email)

	var retrivedPassword string
	err := row.Scan(&u.ID, &retrivedPassword)

	
	//test
	fmt.Printf(" after : u.Email %v",u.Email)
	fmt.Printf(" retrived Password %v",retrivedPassword)
	fmt.Printf("u.ID %v",u.ID)

	if err != nil {
		return errors.New("user not found")
	}

	if u.Password != retrivedPassword {
		return errors.New("password invalid")
	}

	return nil

}

func Save(user models.User) error {
	query := "INSERT INTO users(email,password) VALUES (?,?)"
	//add db
	stmt, err := DB.Prepare(query)
	if err != nil {
		return errors.New("query can't be prepared")
	}
	defer stmt.Close()
	// hashedpass, err := utils.HashedPassword(user.Password)
	// result, err := stmt.Exec(user.Email, hashedpass)
	// if err != nil {
	// 	return err
	// }

	_, err = stmt.Exec(user.Email, user.Password)
	if err != nil {
		return errors.New("query can't be executed")
	}
	// user.ID, err = result.LastInsertId()
	// if err != nil {
	// 	return err
	// }
	return nil
}
