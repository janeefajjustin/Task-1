package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/janeefajjustin/task-1/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	fmt.Println(token)

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "token unavailable"})
		return
	}

	err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err})
		return
	}

	// context.Set("UserId", userId)
	context.Next()
}
