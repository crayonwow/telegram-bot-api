package tgbotapi

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
)

// BotLogger is an interface that represents the required methods to log data.
//
// Instead of requiring the standard logger, we can just specify the methods we
// use and allow users to pass anything that implements these.
type BotLogger interface {
	Println(v ...any)
	Printf(format string, v ...any)
}

type slogWrapper struct {
	*slog.Logger
}

func (sl *slogWrapper) Println(v ...any) {
	sl.Logger.Info(fmt.Sprintln(v...))
}

func (sl *slogWrapper) Printf(format string, v ...any) {
	sl.Logger.Info(fmt.Sprintf(format, v...))
}

var log BotLogger = &slogWrapper{
	slog.New(
		slog.NewTextHandler(
			os.Stderr,
			&slog.HandlerOptions{AddSource: false, Level: slog.LevelInfo},
		),
	),
}

// SetLogger specifies the logger that the package should use.
func SetLogger(logger BotLogger) error {
	if logger == nil {
		return errors.New("logger is nil")
	}
	log = logger
	return nil
}
