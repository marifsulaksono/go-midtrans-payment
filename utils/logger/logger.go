package logger

import (
	"log"
	"os"
)

func OpenFileLogger(path string) (*os.File, error) {
	// load the logger file and create new file if doesn't exist
	loggerFile, err := os.OpenFile(path, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	// add log to file
	log.SetOutput(loggerFile)
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	return loggerFile, nil
}
