package repo

import (
	"errors"

	"github.com/janeefajjustin/task-1/db"
	"github.com/janeefajjustin/task-1/models"
	"github.com/janeefajjustin/task-1/utils"
)

// type UserRepo struct {
// 	db *sql.DB
// }

// func NewUserRepo() UserRepo {
// 	return UserRepo{
// 		db: db.DB,
// 	}
// }

// type RepoInterface interface {
// 	ValidateCredentials(u models.User) error
// }

//Code from repo

func ValidateCredentials(u *models.User) (string,error) {
	query := "SELECT userid, password FROM users WHERE email=$1"

	//test
	// fmt.Printf("u.Email %v", u.Email)
	// fmt.Printf("u.Password %v", u.Password)
	// fmt.Printf("u.ID %vb\n", u.ID)

	row := db.DB.QueryRow(query, u.Email)

	var retrivedPassword string
	err := row.Scan(&u.ID, &retrivedPassword)

	//test
	// fmt.Printf(" after : u.Email %v", u.Email)
	// fmt.Printf(" retrived Password %v", retrivedPassword)
	// fmt.Printf("u.ID %v", u.ID)

	if err != nil {
		return "",errors.New("user not found")
	}

	// if utils.CheckPasswordHash(retrivedPassword, u.Password) != true {
	// 	return "",errors.New("password invalid")
	// }

	return  retrivedPassword,nil

}

func Save(user models.User) error {
	query := "INSERT INTO users(email,password) VALUES ($1,$2)"
	//add db
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return errors.New("query can't be prepared")
	}
	defer stmt.Close()
	hashedpass, err := utils.HashedPassword(user.Password)
	_, err = stmt.Exec(user.Email, hashedpass)
	if err != nil {
		return errors.New("query can't be executed")
	}

	// _, err = stmt.Exec(user.Email, user.Password)
	// if err != nil {
	// 	return errors.New("query can't be executed")
	// }
	// user.ID, err = result.LastInsertId()
	// //fmt.Println(user.ID)
	// if err != nil {
	// 	return err
	// }
	return nil
}
