// sqs_consumer.go
package main

import (
	"fmt"
	"os"
	"encoding/json"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Set the URL of the SQS queue
	queueURL := "http://sqs.us-east-1.localhost.localstack.cloud:4566/000000000000/main-queue"

	// Create a signal channel to handle graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// Start the message consumer in a separate goroutine
	go startMessageConsumer(queueURL, sigCh)

	// Wait for a signal to terminate the program
	<-sigCh
	fmt.Println("Received termination signal. Shutting down...")
}

func startMessageConsumer(queueURL string, sigCh <-chan os.Signal) {
	for {
		select {
		case <-sigCh:
			// Terminate the goroutine on receiving a signal
			return
		default:
			// Receive a message
			output, err := exec.Command("awslocal", "sqs", "receive-message",
				"--queue-url", queueURL,
				"--max-number-of-messages", "1",
				"--wait-time-seconds", "1",
			).CombinedOutput()

			if err != nil {
				fmt.Println("Error receiving message:", err)
				continue
			}

			// Check if there are messages received
			if len(output) > 0 {
				// Process received messages
				fmt.Printf("Received message: %s\n", output)

				// Extract the receipt handle from the received message
				receiptHandle, err := extractReceiptHandle(output)
				if err != nil {
					fmt.Println("Error extracting receipt handle:", err)
					continue
				}

				// Delete the message from the queue after processing
				deleteOutput, err := exec.Command("awslocal", "sqs", "delete-message",
					"--queue-url", queueURL,
					"--receipt-handle", receiptHandle,
				).CombinedOutput()

				if err != nil {
					fmt.Printf("Error deleting message: %v\nCommand output: %s\n", err, deleteOutput)
				} else {
					fmt.Println("Message deleted successfully")
				}
			}

			// Introduce a delay before the next poll
			time.Sleep(1 * time.Second)
		}
	}
}

func extractReceiptHandle(output []byte) (string, error) {
	var message map[string][]map[string]interface{}
	if err := json.Unmarshal(output, &message); err != nil {
		return "", err
	}

	if len(message["Messages"]) == 0 {
		return "", fmt.Errorf("no messages found in the received data")
	}

	receiptHandle, ok := message["Messages"][0]["ReceiptHandle"].(string)
	if !ok {
		return "", fmt.Errorf("ReceiptHandle not found in the received message")
	}

	return receiptHandle, nil
}
