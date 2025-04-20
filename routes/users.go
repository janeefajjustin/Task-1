package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/janeefajjustin/task-1/controller"
)

func Login(context *gin.Context) {
	context.HTML(200, "templates/index.html", gin.H{"title": "Home Page"})
	controller.PostLogin(context)

}

// func Login(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPost {
// 		controller.PostLogin(w, r)
// 	} else {
// 		http.ServeFile(w, r, "templates/login.html")
// 	}

// }

// func Signup(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPost {
// 		controller.PostSignup(w, r)
// 	} else {
// 		http.ServeFile(w, r, "templates/signup.html")
// 	}
// }
