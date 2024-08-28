package config

import (
	"os"
	"path/filepath"

	"log"
)

type Config struct {
	AppDataPath      string
	CachePath        string
	LogsPath         string
	RepoDataFilePath string
}

func (c *Config) Init() {
	path, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("🚨 Error retrieving current user config location: %v", err)
	}
	c.AppDataPath = filepath.Join(path, "incantation")
	c.CachePath = filepath.Join(c.AppDataPath, "cache")
	c.LogsPath = filepath.Join(c.AppDataPath, "logs")

	c.RepoDataFilePath = filepath.Join(c.CachePath, "repos.json")
}
