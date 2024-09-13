package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/andyroid0/incantation/config"
	"github.com/charmbracelet/log"
)

type Logger struct {
	Do log.Logger
}

func (l Logger) ConsoleLogf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func Init(config *config.Config) (*Logger, *os.File) {
	now := time.Now()

	path := fmt.Sprintf(
		"%s%v-%v-%v%s",
		config.LogsPath+"/",
		now.Month(),
		now.Day(),
		now.Year(),
		".log")

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalf("🚨 Error opening file for logger: %v", err)
	}

	charmLog := log.New(file)
	logger := &Logger{Do: *charmLog}

	files, readErr := os.ReadDir(config.LogsPath)
	if readErr != nil {
		log.Fatalf("🚨 Error opening directory to clean old logs: %v", readErr)
	}

	currentTime := time.Now()

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".log") {
			date := strings.TrimSuffix(file.Name(), ".log")
			fileDate, err := time.Parse("January-2-2006", date)
			if err != nil {
				logger.Do.Errorf("🚨 Error parsing date for file:\n\t %s\n\t %s", file.Name(), err)
				continue
			}
			if currentTime.Sub(fileDate).Hours() > 14*24 {
				filePath := filepath.Join(config.LogsPath, file.Name())
				err := os.Remove(filePath)
				if err != nil {
					logger.Do.Errorf("🚨 Error deleting file:\n\t%s\n\t%s", file.Name(), err)
				}
			}

		}
	}

	return logger, file
}
