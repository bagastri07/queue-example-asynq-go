package model

import "context"

type EmailService interface {
	SendEmail(ctx context.Context, email string) error
}
