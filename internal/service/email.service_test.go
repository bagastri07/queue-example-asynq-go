package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEmailService(t *testing.T) {
	tests := []struct {
		name string
		want *EmailService
	}{
		{
			name: "normal",
			want: &EmailService{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewEmailService()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestEmailService_SendEmail(t *testing.T) {
	type args struct {
		ctx   context.Context
		email string
	}
	tests := []struct {
		name        string
		args        args
		expectedErr error
	}{
		{
			name: "normal",
			args: args{
				ctx:   context.Background(),
				email: "email.com",
			},
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EmailService{}

			err := s.SendEmail(tt.args.ctx, tt.args.email)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
