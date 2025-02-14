package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// NewLogger creates a new logger with the provided log level
func NewLogger(level string) zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if level == "debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	return log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}
