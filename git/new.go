package git

import "github.com/andyroid0/incantation/logger"

func New(path string, logger *logger.Logger) Model {
	m := Model{
		Path:   path,
		Logger: logger,
	}
	m.Repository, _ = m.PlainOpen()
	m.Tree, _ = m.WorkTree()
	return m
}
