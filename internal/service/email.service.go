package service

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type EmailService struct{}

func NewEmailService() *EmailService {
	return &EmailService{}
}

func (s *EmailService) SendEmail(ctx context.Context, email string) error {
	// Simulate sending an email
	time.Sleep(3 * time.Second)
	logrus.Info("Email sent:", email)
	return nil
}
