package main

import (
	"net/http"

	"github.com/arjunstein/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main()  {
	server := gin.Default()

	// Method endpoint
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context)  {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"message":"Could not parse request data."})
		return
	}
	
	event.Id = 1
	event.UserId = 1
	context.JSON(http.StatusCreated, gin.H{"message": "Event successfully created", "event": event})
}