package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hey there! well come to my application")

	server := gin.Default()

	server.GET("/shreesh", printShreesh)

	server.Run(":8080")
}

func printShreesh(context *gin.Context) {
	context.JSON(200, "Hey Shreesh Keep Hustling")
}
