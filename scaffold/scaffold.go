package scaffold

import (
	"os"

	"log"

	"github.com/andyroid0/incantation/config"
)

func Go(config *config.Config) {
	mkDirErr := os.Mkdir(config.AppDataPath, 0777)
	if mkDirErr != nil && !os.IsExist(mkDirErr) {
		log.Fatalf("🚨 Error creating AppData directory: %v", mkDirErr)
	}

	mkLogsDirErr := os.Mkdir(config.LogsPath, 0777)
	if mkLogsDirErr != nil && !os.IsExist(mkLogsDirErr) {
		log.Fatalf("🚨 Error creating logs directory for logger: %v", mkLogsDirErr)
	}

	mkCacheDirErr := os.Mkdir(config.CachePath, 0777)
	if mkCacheDirErr != nil && !os.IsExist(mkCacheDirErr) {
		log.Fatalf("🚨 Error creating logs directory for logger: %v", mkCacheDirErr)
	}
}
