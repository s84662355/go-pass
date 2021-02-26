package log

import (
    "github.com/s84662355/goLogger"
    "go.uber.org/zap"
)

var Logger = new(logger)

type logger struct {
    log *zap.Logger
}

func (l *logger) Init(path string, isDay bool) {
    l.log = goLogger.Logger(path, isDay)
}

func (l *logger) Info(msg string) {
    l.log.Info(msg)
}

func (l *logger) Warn(msg string) {
    l.log.Warn(msg)
}

func (l *logger) Error(msg string) {
    l.log.Error(msg)
}

func (l *logger) Debug(msg string) {
    l.log.Debug(msg)
}

func (l *logger) DPanic(msg string) {
    l.log.DPanic(msg)
}

func (l *logger) Panic(msg string) {
    l.log.Panic(msg)
}

func (l *logger) Fatal(msg string) {
    l.log.Fatal(msg)
}
