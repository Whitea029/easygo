package logging

import (
	"{{ .GoModule }}/conf"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapSinLogger *zap.SugaredLogger

type zapLogger struct {
	logConf *conf.Logger
	logger  *zap.SugaredLogger
}

var zapLogLevelMapping = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.ErrorLevel,
}

func newZapLogger() *zapLogger {
	logger := &zapLogger{logConf: &conf.GetConfig().Logger}
	logger.Init()
	return logger
}

func (l *zapLogger) getLogLevel() zapcore.Level {
	level, exists := zapLogLevelMapping[l.logConf.Logger]
	if !exists {
		return zapcore.DebugLevel
	}
	return level
}

func (l *zapLogger) Init() {
	once.Do(func() {
		fileName := fmt.Sprintf("%s%s-%s.%s", l.logConf.FilePath, time.Now().Format("2006-01-02"), uuid.New(), "log")
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   fileName,
			MaxSize:    1,
			MaxAge:     20,
			LocalTime:  true,
			MaxBackups: 5,
			Compress:   true,
		})

		conf := zap.NewProductionEncoderConfig()
		conf.EncodeTime = zapcore.ISO8601TimeEncoder

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(conf),
			w,
			l.getLogLevel())

		logger := zap.New(core, zap.AddCaller(),
			zap.AddCallerSkip(1),
			zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()

		logger = logger.With("AppName", "MyApp", "LoggerName", "Zaplog")

		zapSinLogger = logger
	})

	l.logger = zapSinLogger
}

func (l *zapLogger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogInfo(cat, sub, extra)
	l.logger.Debugw(msg, params...)
}

func (l *zapLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args)
}

func (l *zapLogger) Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogInfo(cat, sub, extra)
	l.logger.Infow(msg, params...)
}

func (l *zapLogger) Infof(templates string, args ...interface{}) {
	l.logger.Infof(templates, args)
}

func (l *zapLogger) Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogInfo(cat, sub, extra)
	l.logger.Warnw(msg, params...)
}

func (l *zapLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args)
}

func (l *zapLogger) Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogInfo(cat, sub, extra)
	l.logger.Errorw(msg, params...)
}

func (l *zapLogger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args)
}

func (l *zapLogger) Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogInfo(cat, sub, extra)
	l.logger.Fatalw(msg, params...)
}

func (l *zapLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args)
}

func prepareLogInfo(cat Category, sub SubCategory, extra map[ExtraKey]interface{}) []interface{} {
	if extra == nil {
		extra = make(map[ExtraKey]interface{})
	}
	extra["Category"] = cat
	extra["SubCategory"] = sub

	return logParamsToZapParams(extra)
}
