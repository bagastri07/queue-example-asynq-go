package main

import (
	"github.com/bagastri07/queue-example-asynq-go/cmd"
	"github.com/bagastri07/queue-example-asynq-go/utils/logger"
)

func main() {
	logger.SetUpLogger()
	cmd.Execute()
}
