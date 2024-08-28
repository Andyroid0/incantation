package logger

import (
	"fmt"
	"os"
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

	return logger, file
}
