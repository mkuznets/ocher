// Package zerologadapter provides a logger that writes to a github.com/rs/zerolog.
package zerologadapter

import (
	"context"

	"github.com/rs/zerolog"
	"mkuznets.com/go/ocher"
)

type Logger struct {
	logger zerolog.Logger
}

// NewLogger accepts a zerolog.Logger as input and returns a new custom pgx
// logging fascade as output.
func NewLogger(logger zerolog.Logger) *Logger {
	return &Logger{logger: logger}
}

func (pl *Logger) Log(ctx context.Context, level ocher.LogLevel, msg string, data map[string]interface{}) {
	var zlevel zerolog.Level
	switch level {
	case ocher.LogLevelNone:
		zlevel = zerolog.NoLevel
	case ocher.LogLevelError:
		zlevel = zerolog.ErrorLevel
	case ocher.LogLevelWarn:
		zlevel = zerolog.WarnLevel
	case ocher.LogLevelInfo:
		zlevel = zerolog.InfoLevel
	case ocher.LogLevelDebug:
		zlevel = zerolog.DebugLevel
	default:
		zlevel = zerolog.DebugLevel
	}

	log := pl.logger.With().Fields(data).Logger()
	log.WithLevel(zlevel).Msg(msg)
}
