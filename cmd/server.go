package cmd

import (
	"time"

	"github.com/bagastri07/queue-example-asynq-go/internal/config"
	"github.com/bagastri07/queue-example-asynq-go/internal/handler"
	"github.com/bagastri07/queue-example-asynq-go/internal/service"
	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func runServer(cmd *cobra.Command, args []string) {
	// Create a new Asynq client
	redisConn := asynq.RedisClientOpt{
		Addr:     config.RedisHost(),
		Password: config.RedisPassword(),
	}

	server := asynq.NewServer(
		redisConn,
		asynq.Config{
			Concurrency: 10,
			Queues:      map[string]int{"default": 1},
		},
	)

	emailSvc := service.NewEmailService()

	handler := handler.NewHandler(emailSvc)

	mux := asynq.NewServeMux()
	mux.HandleFunc("send-email", handler.SendEmailHandler)
	mux.HandleFunc("cron-salutation", handler.CronSalutationHandler)

	go func() {
		scheduler := asynq.NewScheduler(
			redisConn,
			&asynq.SchedulerOpts{},
		)

		opts := []asynq.Option{
			asynq.Retention(30 * time.Second),
			asynq.MaxRetry(10),
		}

		scheduler.Register("@every 5s", asynq.NewTask("cron-salutation", []byte("hello world!!!"), opts...))

		if err := scheduler.Run(); err != nil {
			logrus.Error("Scheduler error:", err)
		}
	}()

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
