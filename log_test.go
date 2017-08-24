package logging

import (
    "testing"
    "os"
)

type LogTest struct {
    logger     ILogger
    showSQL    bool
}

func Test_SimpleLogger(t *testing.T)  {
    l := new(LogTest)
    logger := NewSimpleLogger(os.Stdout)
    l.SetLogger(logger)
    l.ShowSQL(true)

    l.Info("you want to print out message")
}

func (log4j *LogTest) ShowSQL(show... bool) {
    log4j.logger.ShowSQL(show...)
    if len(show) == 0 {
        log4j.showSQL = true
    } else {
        log4j.showSQL = show[0]
    }
}

func (log4j *LogTest) SetLogger(logger ILogger) {
    log4j.logger = logger
}

func (log4j *LogTest) Info(args ...interface{})  {
    log4j.logger.Info(args)
}

func (log4j *LogTest) Debug(args ...interface{})  {
    log4j.logger.Debug(args)
}

func (log4j *LogTest) Warn(args ...interface{})  {
    log4j.logger.Warn(args)
}

func (log4j *LogTest) Error(args ...interface{})  {
    log4j.logger.Error(args)
}