package routes

import (
	"github.com/janeefajjustin/task-1/handler"
	"github.com/janeefajjustin/task-1/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	 h:=handler.UserHandler{}
	 

	// http.HandleFunc("/", Login)
	// http.HandleFunc("/signup", Signup)

	//server.LoadHTMLFiles("templates/login.html")
	server.POST("/user/login", h.Login)
	server.GET("/", handler.LoginPage)
	server.POST("/user/signup", handler.SignUp)
	server.POST("/user/logout", handler.Logout)
	server.GET("/user/:id", middlewares.Authenticate, handler.UserDetails)

}
