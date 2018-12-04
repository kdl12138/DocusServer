package Log

import "os"

func New() *Logger {
	return &Logger{
		Out:   os.Stderr,
		Level: InfoLevel,
	}
}

func (Logger *Logger) Close() {
	Logger.mu.Lock()
}
