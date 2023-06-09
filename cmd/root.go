package cmd

import (
	"os"

	"github.com/bagastri07/queue-example-asynq-go/internal/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "queue-example",
	Short: "A simple queue example",
}

func Execute() {
	config.GetConf()

	if err := rootCmd.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}
