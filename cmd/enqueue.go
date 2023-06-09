package cmd

import (
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

	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379", Password: "mantis"})
	defer client.Close()

	task := asynq.NewTask("send_email", []byte(email))
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
