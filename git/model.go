package git

import (
	"github.com/andyroid0/incantation/logger"

	"github.com/go-git/go-git/v5"
)

type Model struct {
	Path       string // TODO: Should be set and updated by cache
	Repository *git.Repository
	Tree       *git.Worktree
	Logger     logger.Logger
}
