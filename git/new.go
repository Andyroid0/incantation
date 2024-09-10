package git

import "github.com/andyroid0/incantation/logger"

func New(path string, logger *logger.Logger) Model {
	return Model{
		Path:   path,
		Logger: logger,
	}
}
