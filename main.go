package main

import (
	"github.com/gin-gonic/gin"
	"github.com/janeefajjustin/task-1/db"
	"github.com/janeefajjustin/task-1/routes"
)

func main() {
	db.Initialize()
	server := gin.Default()
	//to load all the html pages
	server.LoadHTMLGlob("templates/*")
	routes.RegisterRoutes(server)
	server.Run("localhost:8081")
	//   r.GET("/ping", func(c *gin.Context) {
	//     c.JSON(http.StatusOK, gin.H{
	//       "message": "pong",
	//     })

	// fmt.Println("Server started at http://localhost:8081")
	// if err := http.ListenAndServe(":8081", nil); err != nil {
	// 	fmt.Println("Error starting server:", err)
	// }

}
