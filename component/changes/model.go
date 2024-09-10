package changes

import (
	"github.com/andyroid0/incantation/git"
	"github.com/andyroid0/incantation/logger"
)

type Model struct {
	Logger *logger.Logger
	Git    *git.Model
}
