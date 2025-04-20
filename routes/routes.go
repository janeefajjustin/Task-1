package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	// http.HandleFunc("/", Login)
	// http.HandleFunc("/signup", Signup)

	server.LoadHTMLFiles("templates/login.html")
	server.POST("/", Login)
	//server.GET("/signup", SignUp)

}
