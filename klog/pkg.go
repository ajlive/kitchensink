package klog

import (
	"context"
	"io"
	"log/slog"
	"os"

	"github.com/charmbracelet/lipgloss"
	chlog "github.com/charmbracelet/log"
)

type Formatter chlog.Formatter

const (
	TextFormatter   = Formatter(chlog.TextFormatter)
	JSONFormatter   = Formatter(chlog.JSONFormatter)
	LogfmtFormatter = Formatter(chlog.LogfmtFormatter)
)

const (
	LevelDebug = slog.LevelDebug
	LevelInfo  = slog.LevelInfo
	LevelRun   = slog.Level(2)
	LevelWarn  = slog.LevelWarn
	LevelError = slog.LevelError
)

func New(opts ...option) *slog.Logger {
	handler := newCharmlogger(opts...)
	return slog.New(handler)
}

type CLILogger struct {
	*slog.Logger
}

func NewCLILogger(opts ...option) *CLILogger {
	handler := newCharmlogger(opts...)
	styles := chlog.DefaultStyles()
	styles.Levels[chlog.Level(LevelRun)] = styleLevelRun
	handler.SetStyles(styles)
	return &CLILogger{slog.New(handler)}
}

func (log *CLILogger) Run(msg string) {
	log.Logger.Log(context.Background(), LevelRun, msg)
}

func WithLevel(level slog.Level) option {
	return func(cfg *config) {
		cfg.Level = chlog.Level(level)
	}
}

func WithWriter(w io.Writer) option {
	return func(cfg *config) {
		cfg.writer = w
	}
}

func WithFormatter(formatter Formatter) option {
	return func(cfg *config) {
		cfg.Formatter = chlog.Formatter(formatter)
	}
}

type option func(*config)

type config struct {
	*chlog.Options
	writer io.Writer
}

var styleLevelRun = lipgloss.NewStyle().
	SetString("RUN").
	Bold(true).
	MaxWidth(4).
	Foreground(lipgloss.Color("13"))

func newCharmlogger(opts ...option) *chlog.Logger {
	cfg := &config{
		writer: os.Stderr,
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return chlog.NewWithOptions(cfg.writer, *cfg.Options)
}
