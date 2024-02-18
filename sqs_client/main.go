package main

import (
	"fmt"
	"os/exec"
)

func pushMessageToSQS(messageBody string, queueURL string) error {
	// Execute awslocal command to send a message to the SQS queue
	cmd := exec.Command("awslocal", "sqs", "send-message",
		"--queue-url", queueURL,
		"--message-body", messageBody)

	// Run the command
	err := cmd.Run()
	if err != nil {
		return err
	}

	fmt.Println("Message sent to SQS successfully")
	return nil
}

func main() {
	defaultMessageBody := "Another request sent to the API"
	defaultQueueURL := "http://sqs.us-east-1.localhost.localstack.cloud:4566/000000000000/main-queue"

	// Call pushMessageToSQS with default values
	err := pushMessageToSQS(defaultMessageBody, defaultQueueURL)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
