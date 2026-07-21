package logging

import (
	"io"
	"os"
	"time"

	loggerinterface "github.com/arthurhzna/inventory-reservation-system/backend/internal/domain/logger"
	"github.com/rs/zerolog"
)

type ZeroLogger struct {
	log *zerolog.Logger
}

func NewZeroLogger(
	level int,
) *ZeroLogger {

	log := zerolog.
		New(
			zerolog.ConsoleWriter{
				Out:          os.Stdout,
				TimeLocation: time.UTC,
				TimeFormat:   zerolog.TimeFormatUnix,
			},
		).
		Level(zerolog.Level(level)).
		With().
		Timestamp().
		Logger()

	log = log.With().
		Caller().
		Int("pid", os.Getpid()).
		Logger()

	return &ZeroLogger{
		log: &log,
	}
}

func (l *ZeroLogger) GetWriter() io.Writer {
	return l.log
}

func (l *ZeroLogger) Printf(
	format string,
	args ...any,
) {
	l.log.Printf(format, args...)
}

func (l *ZeroLogger) Error(args ...any) {
	l.log.Error().
		Msg(argsToString(args...))
}

func (l *ZeroLogger) Errorf(
	format string,
	args ...any,
) {
	l.log.Error().
		Msgf(format, args...)
}

func (l *ZeroLogger) Fatal(args ...any) {
	l.log.Fatal().
		Msg(argsToString(args...))
}

func (l *ZeroLogger) Fatalf(
	format string,
	args ...any,
) {
	l.log.Fatal().
		Msgf(format, args...)
}

func (l *ZeroLogger) Info(args ...any) {
	l.log.Info().
		Msg(argsToString(args...))
}

func (l *ZeroLogger) Infof(
	format string,
	args ...any,
) {
	l.log.Info().
		Msgf(format, args...)
}

func (l *ZeroLogger) Warn(args ...any) {
	l.log.Warn().
		Msg(argsToString(args...))
}

func (l *ZeroLogger) Warnf(
	format string,
	args ...any,
) {
	l.log.Warn().
		Msgf(format, args...)
}

func (l *ZeroLogger) Debug(args ...any) {
	l.log.Debug().
		Msg(argsToString(args...))
}

func (l *ZeroLogger) Debugf(
	format string,
	args ...any,
) {
	l.log.Debug().
		Msgf(format, args...)
}

func (l *ZeroLogger) WithField(key string, value any) loggerinterface.Logger {
	var log zerolog.Logger
	if err, ok := value.(error); ok {
		log = l.log.With().Err(err).Logger()
	} else {
		log = l.log.With().Any(key, value).Logger()
	}

	return &ZeroLogger{
		log: &log,
	}
}

func (l *ZeroLogger) WithFields(fields map[string]any) loggerinterface.Logger {
	logCtx := l.log.With()
	for k, v := range fields {
		if errs, ok := v.([]error); ok {
			logCtx = logCtx.Errs(k, errs)
		} else if err, ok := v.(error); ok {
			logCtx = logCtx.AnErr(k, err)
		} else {
			logCtx = logCtx.Any(k, v)
		}
	}

	log := logCtx.Logger()
	return &ZeroLogger{
		log: &log,
	}
}
