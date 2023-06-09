package cmd

import (
	"context"

	"github.com/bagastri07/queue-example-asynq-go/internal/server"
	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func runServer(cmd *cobra.Command, args []string) {
	server := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "localhost:6379", Password: "mantis"},
		asynq.Config{
			Concurrency: 10,
			Queues:      map[string]int{"default": 1},
		},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc("send_email", serverHandler)

	if err := server.Run(mux); err != nil {
		logrus.Error("Server error:", err)
	}
}

func serverHandler(ctx context.Context, task *asynq.Task) error {
	return server.SendEmailHandler(ctx, task)
}

func init() {
	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "Start the server",
		Run:   runServer,
	}

	rootCmd.AddCommand(serverCmd)
}
