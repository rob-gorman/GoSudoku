package logger

import (
	"log"
	"os"
)

type AppLogger struct {
	Inf *log.Logger
	Err *log.Logger
}

func New() *AppLogger {
	return &AppLogger{
		Inf: log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		Err: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *AppLogger) Info(msg string) {
	l.Inf.Println(msg)
}

func (l *AppLogger) Error(msg ...string) {
	l.Err.Println(msg)
}
