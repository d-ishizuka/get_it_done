package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var log zerolog.Logger

func init() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	log = zerolog.New(output).With().Timestamp().Caller().Logger()

	// ログレベル設定
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if os.Getenv("DEBUG") == "true" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}

// GetLogger returns the configured logger
func GetLogger() *zerolog.Logger {
	return &log
}

// Info logs an info message
func Info(msg string, fields ...map[string]interface{}) {
	event := log.Info()
	if len(fields) > 0 {
		for k, v := range fields[0] {
			event = event.Interface(k, v)
		}
	}
	event.Msg(msg)
}

// Error logs an error message
func Error(err error, msg string, fields ...map[string]interface{}) {
		event := log.Error().Err(err)
		if len(fields) > 0 {
			for k, v := range fields[0] {
				event = event.Interface(k, v)
			}
		}
		event.Msg(msg)
}