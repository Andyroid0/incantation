package cache

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/andyroid0/incantation/config"
	"github.com/andyroid0/incantation/logger"
)

type Cache struct {
	Repos  Repos
	Logger *logger.Logger
	Config *config.Config
}

func (c Cache) Init() {
	jsonFile, err := os.Open(c.Config.RepoDataFilePath)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	var repoData Repos
	unMarshalErr := json.Unmarshal(byteValue, &repoData)
	if unMarshalErr != nil {
		c.Logger.Do.Errorf("Error unmarshaling repo data: \n\t%v", unMarshalErr)
		return
	}

	// c.Logger.Do.Print(repoData.Data[0].Name)
	//TODO: do something with the data.
}
