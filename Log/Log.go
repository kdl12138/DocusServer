package Log

import (
	"io"
	"sync"
)

type Logger struct {
	// log target
	Out io.Writer
	// log level
	Level Level
	// Mutex
	mu sync.Mutex
}

func (Logger *Logger) SetLevel(lvl Level) {
	Logger.Level = lvl
}

/*
Info level writer
*/
func (Logger *Logger) Infoln(args ...interface{}) {

}

func (Logger *Logger) Info(args ...interface{}) {

}

func (Logger *Logger) Infof(args ...interface{}) {

}

/*
Debug level writer
*/
func (Logger *Logger) Debugln(args ...interface{}) {

}

func (Logger *Logger) Debug(args ...interface{}) {

}

func (Logger *Logger) Debugf(args ...interface{}) {

}

/*
Warn level writer
*/
func (Logger *Logger) Warnln(args ...interface{}) {

}

func (Logger *Logger) Warn(args ...interface{}) {

}

func (Logger *Logger) Warnf(args ...interface{}) {

}

/*
Error level writer
*/
func (Logger *Logger) Errorln(args ...interface{}) {

}

func (Logger *Logger) Error(args ...interface{}) {

}

func (Logger *Logger) Errorf(args ...interface{}) {

}

/*
Panic level writer
*/
func (Logger *Logger) Panicln(args ...interface{}) {

}

func (Logger *Logger) Panic(args ...interface{}) {

}

func (Logger *Logger) Panicf(args ...interface{}) {

}
