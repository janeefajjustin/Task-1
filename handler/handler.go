package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/janeefajjustin/task-1/models"
	"github.com/janeefajjustin/task-1/repo"
)

func Login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	//testttt
	fmt.Println(user)
	fmt.Printf("user.Email %v \n", user.Email)
	fmt.Printf("user.Password %v \n", user.Password)
	fmt.Printf("user.ID %v \n", user.ID)

	err = repo.ValidateCredentials(&user)
	//err = nil

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"messege": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login Successfull!"})
}

func SignUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request"})
		return
	}

	err = repo.Save(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created !"})

}
