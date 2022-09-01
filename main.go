package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		log.Println("connectivity test")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to You!",
		})
	})

	r.Run(":8877")
}
