package git

import (
	"errors"

	"github.com/andyroid0/incantation/logger"
	"github.com/go-git/go-git/v5"
)

func PlainOpen(path string, logger logger.Logger) (*git.Repository, error) {
	r, err := git.PlainOpen(".")
	if errors.Is(err, git.ErrRepositoryNotExists) {
		logger.Do.Errorf("🚨 Error, couldn't open repo. It doesn't exist:\n\t%v", err)
	} else if err != nil {
		logger.Do.Errorf("🚨 Error, couldn't open repo:\n\t%v", err)
	}
	return r, err
}

func WorkTree(repo *git.Repository, logger logger.Logger) (*git.Worktree, error) {
	tree, err := repo.Worktree()
	if err != nil {
		logger.Do.Errorf("🚨 Couldn't get Worktree:\n\t%v", err)
	}
	return tree, err
}

func GetStatus(tree *git.Worktree, logger logger.Logger) (git.Status, error) {
	status, err := tree.Status()
	if err != nil {
		logger.Do.Errorf("🚨 Couldn't get status:\n\t%v", err)
	}
	return status, err
}
