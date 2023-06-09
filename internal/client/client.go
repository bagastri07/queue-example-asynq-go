package client

import (
	"fmt"

	"github.com/hibiken/asynq"
)

func EnqueueTask(client *asynq.Client, email string) error {
	task := asynq.NewTask("send_email", []byte(email))
	_, err := client.Enqueue(task)
	if err != nil {
		return fmt.Errorf("failed to enqueue task: %w", err)
	}
	return nil
}
