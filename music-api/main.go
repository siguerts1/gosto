// main.go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"music-api/handlers"
	"music-api/sqsclient"
)

func main() {
	// Initialize the Gin router
	router := gin.Default()

	// Define your album-related routes
	router.GET("/albums", func(c *gin.Context) {
		// Trigger SQS message sending before handling the request
		triggerSQSMessage("GET request received")
		handlers.GetAlbums(c)
	})

	router.GET("/albums/:id", func(c *gin.Context) {
		// Trigger SQS message sending before handling the request
		triggerSQSMessage("GET request received")
		handlers.GetAlbumByID(c)
	})

	router.POST("/albums", func(c *gin.Context) {
		// Trigger SQS message sending before handling the request
		triggerSQSMessage("POST request received")
		handlers.PostAlbums(c)
	})

	// Run the server on localhost:8080
	router.Run("0.0.0.0:8080")
}

func triggerSQSMessage(action string) {
	// Customize the message body and queue URL as needed
	messageBody := fmt.Sprintf("API action: %s", action)
	queueURL := "http://sqs.us-east-1.localhost.localstack.cloud:4566/000000000000/main-queue"

	// Send message to SQS
	err := sqsclient.PushMessageToSQS(messageBody, queueURL)
	if err != nil {
		fmt.Println("Error sending message to SQS:", err)
	}
}
