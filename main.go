package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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

	go func() {
		log.Printf("Listening and serving HTTP on [%s] \n", "8877")
		if err := r.Run(":8877"); err != nil {
			log.Fatal(err.Error())
		}
	}()

	// 监听退出消息
	sg := make(chan os.Signal)
	signal.Notify(sg, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGTERM)
	s := <-sg

	log.Println("got signal: ", s, ", shutting down...")

}
