package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"shreesh.com.in/models"
)

func main() {

	fmt.Println("Hey there! well come to my application")

	server := gin.Default()

	server.GET("/events", getAllEvents)
	server.POST("/create/event", createEvent)

	server.Run(":8080")
}

func getAllEvents(c *gin.Context) {
	events := models.GetAllEvents()
	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	var eventObj models.Event
	c.ShouldBindJSON(&eventObj)
	eventObj.Save()
	c.JSON(200, eventObj)
}
