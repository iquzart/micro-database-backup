package logger

import (
	"log/slog"
	"os"
)

type Logger = slog.Logger

func InitLogger() *Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}
