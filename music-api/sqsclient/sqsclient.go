// sqsclient/sqsclient.go
package sqsclient

import (
   "fmt"
   "os/exec"
)

// PushMessageToSQS sends a message to the SQS queue.
func PushMessageToSQS(messageBody string, queueURL string) error {
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
