package routes

import (
	"github.com/janeefajjustin/task-1/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	// http.HandleFunc("/", Login)
	// http.HandleFunc("/signup", Signup)

	//server.LoadHTMLFiles("templates/login.html")
	server.POST("/user/login", handler.Login)
	//server.POST("/signup", SignUp)

}
