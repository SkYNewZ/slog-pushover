package slogpushover_test

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/gregdel/pushover"

	slogpushover "github.com/SkYNewZ/slog-pushover"
)

func Example() {
	handler := slogpushover.NewHandler(&slogpushover.Options{
		Level:       slog.LevelDebug,
		Token:       os.Getenv("PUSHOVER_TOKEN"),
		Recipient:   os.Getenv("PUSHOVER_RECIPIENT"),
		Message:     nil,
		Converter:   nil,
		AddSource:   true,
		ReplaceAttr: nil,
	})

	logger := slog.New(handler)
	logger = logger.With("release", "v1.0.0")

	logger.
		With(
			slog.Group("user",
				slog.String("id", "user-123"),
				slog.Time("created_at", time.Now().AddDate(0, 0, -1)),
			),
		).
		With("environment", "dev").
		With("error", fmt.Errorf("an error")).
		Error("A message")

	// Output:
}

// nolint:govet
func ExampleTitle() {
	handler := slogpushover.NewHandler(&slogpushover.Options{
		Level:       slog.LevelDebug,
		Token:       os.Getenv("PUSHOVER_TOKEN"),
		Recipient:   os.Getenv("PUSHOVER_RECIPIENT"),
		Message:     &pushover.Message{Title: "My App"},
		Converter:   nil,
		AddSource:   true,
		ReplaceAttr: nil,
	})

	logger := slog.New(handler)
	logger = logger.With("release", "v1.0.0")

	logger.
		With(
			slog.Group("user",
				slog.String("id", "user-123"),
				slog.Time("created_at", time.Now().AddDate(0, 0, -1)),
			),
		).
		With("environment", "dev").
		With("error", fmt.Errorf("an error")).
		Error("A message")

	// Output:
}
