package handler

import (
	"context"

	"github.com/bagastri07/queue-example-asynq-go/internal/model"
	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	emailSvc model.EmailService
}

func NewHandler(e model.EmailService) *Handler {
	return &Handler{
		emailSvc: e,
	}
}

func (h *Handler) SendEmailHandler(ctx context.Context, t *asynq.Task) error {
	err := h.emailSvc.SendEmail(ctx, string(t.Payload()))
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
