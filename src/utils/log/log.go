package log

import (
	"context"

	"log"
)

type Config struct {
	Level string
}

type Interface interface {
	Error(ctx context.Context, format string)
	Info(ctx context.Context, format string)
	Fatal(ctx context.Context, format string)
}

// A Logger belongs to the infrastructure layer.
type Logger struct{}

// Init returns a Logger that implements Interface.
func Init(cfg Config) Interface {
	return &Logger{}
}

// LogError is print messages to log.
func (l *Logger) Info(ctx context.Context, format string) {
	log.SetFlags(log.Ldate | log.Ltime)

	log.Printf(format)
}

// Error logs error messages.
func (l *Logger) Error(ctx context.Context, format string) {
	log.SetFlags(log.Ldate | log.Ltime)

	log.Printf(format)
}

// Error logs fatal messages.
func (l *Logger) Fatal(ctx context.Context, format string) {
	log.SetFlags(log.Ldate | log.Ltime)

	log.Fatal(format)
}
