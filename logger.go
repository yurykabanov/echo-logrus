package echologrus

import (
	"io"

	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

type LoggerAdapter struct {
	*logrus.Logger
}

func NewAdapter(logger *logrus.Logger) *LoggerAdapter {
	return &LoggerAdapter{logger}
}

func (adapter LoggerAdapter) Output() io.Writer {
	return adapter.Logger.Out
}

// Log format is completely determined by logrus
func (adapter LoggerAdapter) Prefix() string {
	return ""
}

// Log format is completely determined by logrus
func (adapter LoggerAdapter) SetPrefix(p string) {
	// do nothing
}

func (adapter LoggerAdapter) Level() log.Lvl {
	switch adapter.Logger.Level {
	case logrus.TraceLevel:
		return log.DEBUG
	case logrus.DebugLevel:
		return log.DEBUG
	case logrus.InfoLevel:
		return log.INFO
	case logrus.WarnLevel:
		return log.WARN
	case logrus.ErrorLevel:
		return log.ERROR

	// There's no _good_ public const for Fatal and Panic levels
	case logrus.FatalLevel:
		return log.ERROR
	case logrus.PanicLevel:
		return log.ERROR
	}

	// Fallback to Info level
	return log.INFO
}

// Ignore level as Logger shouldn't be configured by echo.Echo anyway
func (adapter LoggerAdapter) SetLevel(log.Lvl) {
	// do nothing
}

// Log format is completely determined by logrus
func (adapter LoggerAdapter) SetHeader(string) {
	// do nothing
}

func (adapter LoggerAdapter) Printj(j log.JSON) {
	adapter.Logger.WithFields(logrus.Fields(j)).Print()
}

func (adapter LoggerAdapter) Debugj(j log.JSON) {
	adapter.Logger.WithFields(logrus.Fields(j)).Debug()
}

func (adapter LoggerAdapter) Infoj(j log.JSON) {
	adapter.Logger.WithFields(logrus.Fields(j)).Info()
}

func (adapter LoggerAdapter) Warnj(j log.JSON) {
	adapter.Logger.WithFields(logrus.Fields(j)).Warn()
}

func (adapter LoggerAdapter) Errorj(j log.JSON) {
	adapter.Logger.WithFields(logrus.Fields(j)).Error()
}

func (adapter LoggerAdapter) Fatalj(j log.JSON) {
	adapter.Logger.WithFields(logrus.Fields(j)).Fatal()
}

func (adapter LoggerAdapter) Panicj(j log.JSON) {
	adapter.Logger.WithFields(logrus.Fields(j)).Panic()
}
