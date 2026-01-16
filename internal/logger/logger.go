package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/darkprof83/webdavd/pkg/slogdiscard"
	"github.com/darkprof83/webdavd/pkg/slogpretty"
)

var (
	ErrIncorrectProfile = fmt.Errorf("incorrect logging profile")
)

type Logger struct {
	slog.Logger
}

type profile struct {
	out   io.Writer
	level slog.Level
	setup setupFn
}

type setupFn func(io.Writer, slog.Level) *slog.Logger

var m = map[string]profile{
	"default": {out: os.Stdout, level: slog.LevelDebug, setup: setupText},
	"local":   {out: os.Stdout, level: slog.LevelDebug, setup: setupPretty},
	"dev":     {out: os.Stdout, level: slog.LevelDebug, setup: setupJSON},
	"prod":    {out: os.Stdout, level: slog.LevelInfo, setup: setupText},
	"test":    {out: os.Stdout, level: slog.LevelInfo, setup: setupDiscard},
}

func New() *Logger {
	return &Logger{*setupText(os.Stdout, slog.LevelDebug)}
}

func (log *Logger) Setup(env string) error {
	var p profile
	var ok bool
	if p, ok = m[env]; !ok {
		return ErrIncorrectProfile
	}
	log.Logger = *p.setup(p.out, p.level)
	if p.level == slog.LevelDebug {
		log.Info("debug messages are enabled")
	} else {
		log.Info("debug messages are disabled")
	}
	return nil
}

func setupPretty(out io.Writer, level slog.Level) *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: level,
		},
	}

	handler := opts.NewPrettyHandler(out)

	return slog.New(handler)
}

func setupJSON(out io.Writer, level slog.Level) *slog.Logger {
	return slog.New(slog.NewJSONHandler(out, &slog.HandlerOptions{Level: level}))
}

func setupText(out io.Writer, level slog.Level) *slog.Logger {
	return slog.New(slog.NewTextHandler(out, &slog.HandlerOptions{Level: level}))
}

func setupDiscard(out io.Writer, level slog.Level) *slog.Logger {
	return slogdiscard.NewDiscardLogger()
}
