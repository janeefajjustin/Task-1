package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/janeefajjustin/task-1/db"
)

func PostLogin(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")

	fmt.Printf("Received: Username=%s, Password=%s", username, password)

	err := db.ValidateCredentials(username, password)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Received: Username=%s, Password=%s", username, password)
	}
}

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
