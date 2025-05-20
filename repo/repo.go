package repo

import (
	"database/sql"
	"errors"

	"github.com/janeefajjustin/task-1/db"
	"github.com/janeefajjustin/task-1/models"
	"github.com/janeefajjustin/task-1/utils"
)

type UserRepo struct {
	Db *sql.DB
}

type RepoInterface interface {
	ValidateCredentials(u *models.User) error
}

func NewUserRepo(db *sql.DB) UserRepo {
	return UserRepo{
		Db: db,
	}
}

func(r UserRepo) ValidateCredentials(u *models.User) (string, error) {
	query := "SELECT userid, password FROM users WHERE email=$1"



	row := r.Db.QueryRow(query, u.Email)

	var retrivedPassword string
	err := row.Scan(&u.ID, &retrivedPassword)

	

	if err != nil {
		return "", errors.New("user not found")
	}

	// if utils.CheckPasswordHash(retrivedPassword, u.Password) != true {
	// 	return "",errors.New("password invalid")
	// }

	return retrivedPassword, nil

}

func Save(user models.User) error {
	query := "INSERT INTO users(email,password) VALUES ($1,$2)"

	//var LastInsertId int64
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

	// user.ID, err = result.LastInsertId()
	// fmt.Println(user.ID)

	// if err != nil {
	// 	return err
	// }
	return nil
}

func GetUserDetails(id int64) (models.User, error) {
	query := "SELECT email FROM users WHERE userid=$1"

	row := db.DB.QueryRow(query, id)

	var user models.User
	err := row.Scan(&user.Email)

	user.ID = id

	if err != nil {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}

// func Save(user models.User) error {
// 	query := "INSERT INTO users(email,password) VALUES ($1,$2)"

// 	//var LastInsertId int64
// 	stmt, err := db.DB.Prepare(query)
// 	if err != nil {
// 		return errors.New("query can't be prepared")
// 	}
// 	defer stmt.Close()
// 	hashedpass, err := utils.HashedPassword(user.Password)
// 	_, err = stmt.Exec(user.Email, hashedpass)
// 	if err != nil {
// 		return errors.New("query can't be executed")
// 	}

// 	// _, err = stmt.Exec(user.Email, user.Password)
// 	// if err != nil {
// 	// 	return errors.New("query can't be executed")
// 	// }
// 	// user.ID, err = result.LastInsertId()
// 	// //fmt.Println(user.ID)
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	return nil
// }
