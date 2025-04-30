package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/janeefajjustin/task-1/db"
	"github.com/janeefajjustin/task-1/models"
)

func Login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.ValidateCredentials()
	err = nil

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"messege": "err.Error()"})

		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login Successfull!"})
}

func (u models.User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email=?"
	row := db.DB.QueryRow(query, u.Email)

	var retrivedPassword string
	err := row.Scan(&u.ID, &retrivedPassword)

	if err != nil {
		return errors.New("credentials invalid")
	}

	if u.Password != retrivedPassword {
		return errors.New("credentials invalid")
	}

	return nil

}
