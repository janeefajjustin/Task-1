package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/janeefajjustin/task-1/models"
	"github.com/janeefajjustin/task-1/repo"
	"github.com/janeefajjustin/task-1/service"
	"github.com/janeefajjustin/task-1/utils"
)

type UserHandler struct {
	// UserService *service.UserService
}

// type HandlerInterface interface {
// 	UserDetails(context *gin.Context)
// 	LoginPage(context *gin.Context)
// 	SignUp(context *gin.Context)
// 	Logout(context *gin.Context)
// }

// func NewHandlerService(userService *service.UserService) UserHandler {
// 	return UserHandler{
// 		UserService: userService,
// 	}
// }

func UserDetails(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	user, err := repo.GetUserDetails(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not get the user data"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"user": user})
}

func LoginPage(context *gin.Context) {
	context.HTML(http.StatusOK, "login.html", gin.H{"message": "success"})
	//Login(context)
}

func(h *UserHandler) Login(context *gin.Context) {

	contentType := context.GetHeader("Content-Type")
	var user models.User

	if contentType == "application/json" {
		err := context.ShouldBindJSON(&user)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
			return
		}
	} else {
		user.Email = context.PostForm("username")
		user.Password = context.PostForm("password")
	}

	//actual
	// var user models.User
	// err := context.ShouldBindJSON(&user)

	// if err != nil {
	// 	context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
	// 	return
	// }

	// err := s.CompareUsernameandPassword(&user)
	err := h.UserService.CompareUsernameandPassword(&user)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"messege": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "could not create token"})
	}
	context.SetCookie("jwt-token", token, -1, "/", "localhost", false, true)
	//context.JSON(http.StatusOK, gin.H{"message": "Login Successfull!", "token": token})
	context.HTML(http.StatusOK, "home.html", gin.H{"message": "success"})
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

func Logout(context *gin.Context) {

	// context.SetCookie("session-name", "", -1, "/", "localhost", false, true)

	// Clear the JWT token cookie
	context.SetCookie("jwt-token", "", -1, "/", "localhost", false, true)

	context.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})

}

// func PostLogin(context *gin.Context) {
// 	username := context.PostForm("username")
// 	password := context.PostForm("password")

// 	fmt.Printf("Received: Username=%s, Password=%s", username, password)

// 	err := db.ValidateCredentials(username, password)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Printf("Received: Username=%s, Password=%s", username, password)
// 	}
// }

// func PostLogin(w http.ResponseWriter, r *http.Request) {
// 	username := r.FormValue("username")
// 	password := r.FormValue("password")

// 	fmt.Fprintf(w, "Received: Username=%s, Password=%s", username, password)

// 	err := db.ValidateCredentials(username, password)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Fprintf(w, "Received: Username=%s, Password=%s", username, password)
// 	}
// }

// func PostSignup(w http.ResponseWriter, r *http.Request) {
// 	username := r.FormValue("username")
// 	password := r.FormValue("password")
// 	email := r.FormValue("email")

// 	fmt.Fprintf(w, "Received: Username=%s, Password=%s , Email=%s", username, password, email)
// }
