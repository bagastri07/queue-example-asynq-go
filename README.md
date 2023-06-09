# Asynq Queue Process Example

This is a simple example application that demonstrates how to use Asynq, a powerful asynchronous task processing library, in a Go project.

## Prerequisites

Before running the application, make sure you have the following installed on your system:

- Go (version 1.16 or higher)
- Redis (version 5 or higher)

## Getting Started

Follow these steps to get started with the Asynq Queue Process Example:

1. **Clone the Repository**: Clone this repository to your local machine.

2. **Install Dependencies and Fill Config**: Navigate to the root directory of the project and run the following command to install the required dependencies and dont forget fill the **config.yml** file:

```shell
go mod tidy
```
3. **Start the Server**: In a terminal, run the following command to start the server:
```shell
go run main.go server
```

> This command starts the server and configures it to listen for
> incoming tasks on the "send_email" queue.
4.  **Enqueue a Task**: In another terminal session, run the following command to enqueue a task:
```shell
go run main.go enqueue [email]
```
> Replace `[email]` with the email address to which you want to send the
> email. This command enqueues a "send_email" task with the provided
> email as the payload.
5.  **Monitor the Queue**: To monitor the queue and see the tasks being processed, you can use Asynqmon, a monitoring tool for Asynq. Follow the instructions provided in the [Asynqmon repository](https://github.com/hibiken/asynqmon#release-binaries) to download and start Asynqmon.
Once Asynqmon is running, you should be able to see the tasks being processed in the UI.

That's it! You have successfully set up and used Asynq to process tasks asynchronously in your application.