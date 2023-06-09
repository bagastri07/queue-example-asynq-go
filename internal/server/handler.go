package server

import (
	"context"
	"time"

	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
)

func SendEmailHandler(ctx context.Context, t *asynq.Task) error {
	// Simulate sending an email
	time.Sleep(3 * time.Second)
	logrus.Info("Email sent:", string(t.Payload()))
	return nil
}
