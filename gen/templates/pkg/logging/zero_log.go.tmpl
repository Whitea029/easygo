package logging

import (
	"{{ .GoModule }}/conf"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var zeroSinLogger *zerolog.Logger

type zeroLogger struct {
	logConf *conf.Logger
	logger  *zerolog.Logger
}

var zeroLogLevelMapping = map[string]zerolog.Level{
	"debug": zerolog.DebugLevel,
	"info":  zerolog.InfoLevel,
	"warn":  zerolog.WarnLevel,
	"error": zerolog.ErrorLevel,
	"fatal": zerolog.FatalLevel,
}

func newZeroLogger() *zeroLogger {
	logger := &zeroLogger{logConf: &conf.GetConfig().Logger}
	logger.Init()
	return logger
}

func (l *zeroLogger) getLogLevel() zerolog.Level {
	level, exists := zeroLogLevelMapping[l.logConf.Level]
	if !exists {
		return zerolog.DebugLevel
	}
	return level
}

func (l *zeroLogger) Init() {
	once.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		fileName := fmt.Sprintf("%s%s-%s.%s", l.logConf.FilePath, time.Now().Format("2024-12-19"), uuid.New(), "log")

		file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			panic("could not open file")
		}

		var logger = zerolog.New(file).
			With().
			Timestamp().
			Str("AppName", "MyApp"). // TODO
			Str("LoggerName", "Zerolog").
			Logger()
		zerolog.SetGlobalLevel(l.getLogLevel())

		zeroSinLogger = &logger
	})
	l.logger = zeroSinLogger
}

func (l *zeroLogger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	l.logger.Debug().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(logParamsToZeroParams(extra)).
		Msg(msg)
}

func (l *zeroLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debug().Msgf(template, args)
}

func (l *zeroLogger) Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	l.logger.Info().
		Str("category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(logParamsToZeroParams(extra)).
		Msg(msg)
}

func (l *zeroLogger) Infof(template string, args ...interface{}) {
	l.logger.Info().Msgf(template, args)
}

func (l *zeroLogger) Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	l.logger.Warn().
		Str("category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(logParamsToZeroParams(extra)).
		Msg(msg)
}

func (l *zeroLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warn().Msgf(template, args)
}

func (l *zeroLogger) Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	l.logger.Error().
		Str("category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(logParamsToZeroParams(extra)).
		Msg(msg)
}

func (l *zeroLogger) Errorf(template string, args ...interface{}) {
	l.logger.Error().Msgf(template, args)
}

func (l *zeroLogger) Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	l.logger.Fatal().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(logParamsToZeroParams(extra)).
		Msg(msg)
}

func (l *zeroLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatal().Msgf(template, args)
}
