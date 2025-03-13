package logger

import (
	"log"
	"time"
)

// Logger - структура логгера
type Logger struct {
	prefix string
}

// NewLogger - конструктор для Logger
func NewLogger(prefix string) *Logger {
	return &Logger{prefix: prefix}
}

// Log - метод для логирования сообщений
func (l *Logger) Log(message string) {
	log.Printf("[%s] %s: %s\n", time.Now().Format("2006-01-02 15:04:05"), l.prefix, message)
}
