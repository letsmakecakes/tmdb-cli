package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Create a context that can be canceled
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle application startup and graceful shutdown
	if err := run(ctx); err != nil {
		log.Errorf("Application failed: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	// Setup graceful shutdown
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load configuration: %v", err)
	}

	// Validate configuration
	if err := cfg.Validate(); err != nil {
		return fmt.Errorf("invalid configuration: %v", err)
	}

	// Create API client with timeout and retry
	apiClient, err := api.NewClient(
		cfg.APIKey,
		api.WithTimeout(10*time.Second),
		api.WithRetryCount(3),
	)
	if err != nil {
		return fmt.Errorf("failed to create API client: %w", err)
	}

	// Create CLI command
	command, err := cli.NewRootCommand(
		cli.WithConfig(cfg),
		cli.WithAPIClient(apiClient))
	if err != nil {
		return fmt.Errorf("failed to create CLI command: %v", err)
	}

	// Goroutine for handling graceful shutdown
	go func() {
		<-stopChan
		log.Info("Received shutdown signal")
		cancel()
	}()

	// Execute the command
	if err := command.ExecuteContext(ctx); err != nil {
		return fmt.Errorf("command execution failed: %w", err)
	}

	return nil
}
