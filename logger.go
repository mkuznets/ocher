package ocher

import (
	"context"
)

// The values for log levels are chosen such that the zero value means that no
// log level was specified.
const (
	LogLevelTrace = 6
	LogLevelDebug = 5
	LogLevelInfo  = 4
	LogLevelWarn  = 3
	LogLevelError = 2
	LogLevelNone  = 1
)

// LogLevel represents the pgx logging level. See LogLevel* constants for
// possible values.
type LogLevel int

// Logger is the interface used to get logging from pgx internals.
type Logger interface {
	// Log a message at the given level with data key/value pairs. data may be nil.
	Log(ctx context.Context, level LogLevel, msg string, data map[string]interface{})
}

type nopLogger struct{}

func (*nopLogger) Log(ctx context.Context, level LogLevel, msg string, data map[string]interface{}) {}
