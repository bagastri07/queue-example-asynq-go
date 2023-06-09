package handler

import (
	"context"
	"errors"
	"testing"

	"github.com/bagastri07/queue-example-asynq-go/internal/model"
	"github.com/bagastri07/queue-example-asynq-go/internal/model/mock_model"
	"github.com/golang/mock/gomock"
	"github.com/hibiken/asynq"
	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	var (
		ctrl         = gomock.NewController(t)
		mockEmailSvc = mock_model.NewMockEmailService(ctrl)
	)

	type args struct {
		e model.EmailService
	}
	tests := []struct {
		name string
		args args
		want *Handler
	}{
		{
			name: "normal",
			args: args{
				e: mockEmailSvc,
			},
			want: &Handler{emailSvc: mockEmailSvc},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewHandler(tt.args.e)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestHandler_SendEmailHandler(t *testing.T) {
	var (
		ctrl         = gomock.NewController(t)
		mockEmailSvc = mock_model.NewMockEmailService(ctrl)
	)
	type args struct {
		ctx context.Context
		t   *asynq.Task
	}
	tests := []struct {
		name        string
		args        args
		expectedErr error
		prepareMock func()
	}{
		{
			name: "normal",
			args: args{
				ctx: context.Background(),
				t:   asynq.NewTask("send-email", []byte("email.com")),
			},
			expectedErr: nil,
			prepareMock: func() {
				mockEmailSvc.EXPECT().
					SendEmail(context.Background(), "email.com").
					Times(1).
					Return(nil)
			},
		},
		{
			name: "err",
			args: args{
				ctx: context.Background(),
				t:   asynq.NewTask("send-email", []byte("email.com")),
			},
			expectedErr: errors.New("err"),
			prepareMock: func() {
				mockEmailSvc.EXPECT().
					SendEmail(context.Background(), "email.com").
					Times(1).
					Return(errors.New("err"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareMock()

			h := &Handler{
				emailSvc: mockEmailSvc,
			}

			err := h.SendEmailHandler(tt.args.ctx, tt.args.t)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
