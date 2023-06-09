package cmd

import (
	"github.com/bagastri07/queue-example-asynq-go/internal/config"
	"github.com/bagastri07/queue-example-asynq-go/internal/handler"
	"github.com/bagastri07/queue-example-asynq-go/internal/service"
	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func runServer(cmd *cobra.Command, args []string) {
	server := asynq.NewServer(
		asynq.RedisClientOpt{Addr: config.RedisHost(), Password: config.RedisPassword()},
		asynq.Config{
			Concurrency: 10,
			Queues:      map[string]int{"default": 1},
		},
	)

	emailSvc := service.NewEmailService()

	handler := handler.NewHandler(emailSvc)

	mux := asynq.NewServeMux()
	mux.HandleFunc("send-email", handler.SendEmailHandler)

	if err := server.Run(mux); err != nil {
		logrus.Error("Server error:", err)
	}
}
func init() {
	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "Start the server",
		Run:   runServer,
	}

	rootCmd.AddCommand(serverCmd)
}
