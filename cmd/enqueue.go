package cmd

import (
	"github.com/bagastri07/queue-example-asynq-go/internal/config"
	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func enqueueTask(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		logrus.Info("Please provide the email address as an argument.")
		return
	}

	email := args[0]

	client := asynq.NewClient(asynq.RedisClientOpt{Addr: config.RedisHost(), Password: config.RedisPassword()})
	defer client.Close()

	opts := []asynq.Option{
		asynq.Retention(config.RedisCacheTTL()),
		asynq.MaxRetry(10),
	}

	task := asynq.NewTask("send-email", []byte(email), opts...)
	result, err := client.Enqueue(task)
	if err != nil {
		logrus.Info("Failed to enqueue task:", err)
		return
	}

	logrus.Info("Task enqueued:", result)
}

func init() {
	enqueueCmd := &cobra.Command{
		Use:   "enqueue [email]",
		Short: "Enqueue a task",
		Run:   enqueueTask,
		Args:  cobra.ExactArgs(1), // Expects exactly one argument (email)
	}

	rootCmd.AddCommand(enqueueCmd)
}
