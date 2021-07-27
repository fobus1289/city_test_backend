package service

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

type Logger struct {
	*log.Logger
	once *sync.Once
}

func NewLoggerService(out io.Writer) *Logger {

	logger := &Logger{
		Logger: log.New(out, "INFO: ", log.LstdFlags|log.Lshortfile),
	}

	go logger.file()

	return logger
}

func (l *Logger) file() {

	for {
		filename := fmt.Sprintf("log/%s.%s", time.Now().String()[:10], "log")
		file, _ := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)

		l.Logger.SetOutput(file)
		time.Sleep(24 * time.Hour)
	}
}
