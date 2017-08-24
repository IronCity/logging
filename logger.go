package logging

import (
    "log"
    "io"
    "fmt"
    "runtime"
    "os"
)

const (
    DEFAULT_LOG_PREFIX = ""
    DEFAULT_LOG_FLAG   = log.Ldate | log.Ltime
    DEFAULT_LOG_LEVEL  = LOG_DEBUG
)

var _ILogger = DiscardLogger{}

type DiscardLogger struct{}

// Debug empty implementation
func (DiscardLogger) Debug(v ...interface{}) {}

// Debugf empty implementation
func (DiscardLogger) Debugf(format string, v ...interface{}) {}

// Error empty implementation
func (DiscardLogger) Error(v ...interface{}) {}

// Errorf empty implementation
func (DiscardLogger) Errorf(format string, v ...interface{}) {}

// Info empty implementation
func (DiscardLogger) Info(v ...interface{}) {}

// Infof empty implementation
func (DiscardLogger) Infof(format string, v ...interface{}) {}

// Warn empty implementation
func (DiscardLogger) Warn(v ...interface{}) {}

// Warnf empty implementation
func (DiscardLogger) Warnf(format string, v ...interface{}) {}

// Level empty implementation
func (DiscardLogger) Level() LogLevel {
    return LOG_UNKNOWN
}

// SetLevel empty implementation
func (DiscardLogger) SetLevel(l LogLevel) {}

// ShowSQL empty implementation
func (DiscardLogger) ShowSQL(show ...bool) {}

// IsShowSQL empty implementation
func (DiscardLogger) IsShowSQL() bool {
    return false
}

// SimpleLogger is the default implment of core.ILogger
type SimpleLogger struct {
    DEBUG   *log.Logger
    ERR     *log.Logger
    INFO    *log.Logger
    WARN    *log.Logger
    level   LogLevel
    showSQL bool
}

var _ ILogger = &SimpleLogger{}

// NewSimpleLogger use a special io.Writer as logger output
func NewSimpleLogger(out io.Writer) *SimpleLogger {
    return NewSimpleLogger2(out, DEFAULT_LOG_PREFIX, DEFAULT_LOG_FLAG)
}

// NewSimpleLogger2 let you customrize your logger prefix and flag
func NewSimpleLogger2(out io.Writer, prefix string, flag int) *SimpleLogger {
    return NewSimpleLogger3(out, prefix, flag, DEFAULT_LOG_LEVEL)
}

// NewSimpleLogger3 let you customrize your logger prefix and flag and logLevel
func NewSimpleLogger3(out io.Writer, prefix string, flag int, l LogLevel) *SimpleLogger {
    return &SimpleLogger{
        DEBUG: log.New(out, fmt.Sprintf("%s [debug] ", prefix), flag),
        ERR:   log.New(out, fmt.Sprintf("%s [error] ", prefix), flag),
        INFO:  log.New(out, fmt.Sprintf("%s [info]  ", prefix), flag),
        WARN:  log.New(out, fmt.Sprintf("%s [warn]  ", prefix), flag),
        level: l,
    }
}

func (s *SimpleLogger) getRouter() (string,int) {
    _, file, line, _ := runtime.Caller(3)
    return file,line
}

// Error implement core.ILogger
func (s *SimpleLogger) Error(v ...interface{}) {
    f,l := s.getRouter()
    if s.level <= LOG_ERR {
        s.ERR.Output(2, fmt.Sprintln(f, ":", l, ":", fmt.Sprint(v...)))
    }
    return
}

// Errorf implement core.ILogger
func (s *SimpleLogger) Errorf(format string, v ...interface{}) {
    if s.level <= LOG_ERR {
        s.ERR.Output(2, fmt.Sprintf(format, v...))
    }
    return
}

// Debug implement core.ILogger
func (s *SimpleLogger) Debug(v ...interface{}) {
    f,l := s.getRouter()
    if s.level <= LOG_DEBUG {
        s.DEBUG.Output(2, fmt.Sprintln(f, ":", l, ":", fmt.Sprint(v...)))
    }
    return
}

// Debugf implement core.ILogger
func (s *SimpleLogger) Debugf(format string, v ...interface{}) {
    if s.level <= LOG_DEBUG {
        s.DEBUG.Output(2, fmt.Sprintf(format, v...))
    }
    return
}

// Info implement core.ILogger
func (s *SimpleLogger) Info(v ...interface{}) {
    f,l := s.getRouter()
    if s.level <= LOG_INFO {
        s.INFO.Output(2, fmt.Sprintln(f, ":", l, ":", fmt.Sprint(v...)))
    }
    return
}

// Infof implement core.ILogger
func (s *SimpleLogger) Infof(format string, v ...interface{}) {
    if s.level <= LOG_INFO {
        s.INFO.Output(2, fmt.Sprintf(format, v...))
    }
    return
}

// Warn implement core.ILogger
func (s *SimpleLogger) Warn(v ...interface{}) {
    f,l := s.getRouter()
    if s.level <= LOG_WARNING {
        s.WARN.Output(2, fmt.Sprintln(f, ":", l, ":", fmt.Sprint(v...)))
    }
    return
}

// Warnf implement core.ILogger
func (s *SimpleLogger) Warnf(format string, v ...interface{}) {
    if s.level <= LOG_WARNING {
        s.WARN.Output(2, fmt.Sprintf(format, v...))
    }
    return
}

// Level implement core.ILogger
func (s *SimpleLogger) Level() LogLevel {
    return s.level
}

// SetLevel implement core.ILogger
func (s *SimpleLogger) SetLevel(l LogLevel) {
    s.level = l
    return
}

// ShowSQL implement core.ILogger
func (s *SimpleLogger) ShowSQL(show ...bool) {
    if len(show) == 0 {
        s.showSQL = true
        return
    }
    s.showSQL = show[0]
}

// IsShowSQL implement core.ILogger
func (s *SimpleLogger) IsShowSQL() bool {
    return s.showSQL
}

// SimpleLogger is the default implment of core.ILogger
type FileLogger struct {
    DEBUG   *log.Logger
    ERR     *log.Logger
    INFO    *log.Logger
    WARN    *log.Logger
    level   LogLevel
    showSQL     bool
}

var _ ILogger = &FileLogger{}

func NewFileLogger(file *os.File) *FileLogger {
    return NewFileLogger2(file, DEFAULT_LOG_PREFIX, DEFAULT_LOG_FLAG)
}

func NewFileLogger2(file *os.File, prefix string, flag int) *FileLogger {
    return NewFileLogger3(file, prefix, flag, DEFAULT_LOG_LEVEL)
}

func NewFileLogger3(file *os.File, prefix string, flag int, l LogLevel) *FileLogger {
    return &FileLogger{
        DEBUG: log.New(file, fmt.Sprintf("%s [debug] ", prefix), flag),
        ERR:   log.New(file, fmt.Sprintf("%s [error] ", prefix), flag),
        INFO:  log.New(file, fmt.Sprintf("%s [info]  ", prefix), flag),
        WARN:  log.New(file, fmt.Sprintf("%s [warn]  ", prefix), flag),
        level: l,
    }
}

func (s *FileLogger) getRouter() (string,int) {
    _, file, line, _ := runtime.Caller(3)
    return file,line
}

// Error implement core.ILogger
func (s *FileLogger) Error(v ...interface{}) {
    f,l := s.getRouter()
    if s.level <= LOG_ERR {
        if s.showSQL {
            log.Print(fmt.Sprintln(f, ":", l, ":", fmt.Sprint(v...)))
        }
        s.ERR.Output(2, fmt.Sprintln(f, ":", l, ":", fmt.Sprint(v...)))
    }
    return
}

// Errorf implement core.ILogger
func (s *FileLogger) Errorf(format string, v ...interface{}) {
    if s.level <= LOG_ERR {
        s.ERR.Output(2, fmt.Sprintf(format, v...))
    }
    return
}

// Debug implement core.ILogger
func (s *FileLogger) Debug(v ...interface{}) {
    f,l := s.getRouter()
    if s.level <= LOG_DEBUG {
        if s.showSQL {
            log.Print(fmt.Sprintln(f, ":", l, ":", fmt.Sprint(v...)))
        }
        s.DEBUG.Output(2, fmt.Sprintln(f, ":", l, ":", fmt.Sprint(v...)))
    }
    return
}

// Debugf implement core.ILogger
func (s *FileLogger) Debugf(format string, v ...interface{}) {
    if s.level <= LOG_DEBUG {
        s.DEBUG.Output(2, fmt.Sprintf(format, v...))
    }
    return
}

// Info implement core.ILogger
func (s *FileLogger) Info(v ...interface{}) {
    f,l := s.getRouter()
    if s.level <= LOG_INFO {
        if s.showSQL {
            log.Print(fmt.Sprintln(f, ":", l, ":", fmt.Sprint(v...)))
        }
        s.INFO.Output(2, fmt.Sprintln(f, ":", l, ":", fmt.Sprint(v...)))
    }
    return
}

// Infof implement core.ILogger
func (s *FileLogger) Infof(format string, v ...interface{}) {
    if s.level <= LOG_INFO {
        s.INFO.Output(2, fmt.Sprintf(format, v...))
    }
    return
}

// Warn implement core.ILogger
func (s *FileLogger) Warn(v ...interface{}) {
    f,l := s.getRouter()
    if s.level <= LOG_WARNING {
        if s.showSQL {
            log.Print(fmt.Sprintln(f, ":", l, ":", fmt.Sprint(v...)))
        }
        s.WARN.Output(2, fmt.Sprintln(f, ":", l, ":", fmt.Sprint(v...)))
    }
    return
}

// Warnf implement core.ILogger
func (s *FileLogger) Warnf(format string, v ...interface{}) {
    if s.level <= LOG_WARNING {
        s.WARN.Output(2, fmt.Sprintf(format, v...))
    }
    return
}

// Level implement core.ILogger
func (s *FileLogger) Level() LogLevel {
    return s.level
}

// SetLevel implement core.ILogger
func (s *FileLogger) SetLevel(l LogLevel) {
    s.level = l
    return
}

// ShowSQL implement core.ILogger
func (s *FileLogger) ShowSQL(show ...bool) {
    if len(show) == 0 {
        s.showSQL = true
        return
    }
    s.showSQL = show[0]
}

// IsShowSQL implement core.ILogger
func (s *FileLogger) IsShowSQL() bool {
    return s.showSQL
}

