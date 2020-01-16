package logger

import (
	"log"
	"os"
)

func InitLoggerFile(logpath string) *log.Logger {
	if logpath == "" {
		return log.New(os.Stderr, "", 0)
	}

	file, err := os.OpenFile(logpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	return log.New(file, "", 0)
}
