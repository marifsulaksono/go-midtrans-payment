package logger

import (
	"log"
	"os"
)

func OpenFileErrorLogger(path string) (*os.File, error) {
	loggerFile, err := os.OpenFile(path, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	log.SetOutput(loggerFile)
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	return loggerFile, nil
}

func OpenFileNotifLogger(path string) (*os.File, error) {
	loggerFile, err := os.OpenFile(path, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	log.SetOutput(loggerFile)
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	return loggerFile, nil
}
