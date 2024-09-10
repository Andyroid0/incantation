package git

import (
	"errors"

	"github.com/go-git/go-git/v5"
)

func (m *Model) PlainOpen() (*git.Repository, error) {
	r, err := git.PlainOpen(m.Path)
	if errors.Is(err, git.ErrRepositoryNotExists) {
		m.Logger.Do.Errorf("🚨 Error, couldn't open repo. It doesn't exist:\n\t%v", err)
	} else if err != nil {
		m.Logger.Do.Errorf("🚨 Error, couldn't open repo:\n\t%v", err)
	}
	return r, err
}

func (m *Model) WorkTree() (*git.Worktree, error) {
	if m.Repository == nil {
		m.Repository, _ = m.PlainOpen()
	}
	tree, err := m.Repository.Worktree()
	if err != nil {
		m.Logger.Do.Errorf("🚨 Couldn't get Worktree:\n\t%v", err)
	}
	return tree, err
}

func (m *Model) GetStatus() (git.Status, error) {
	if m.Tree == nil {
		m.Tree, _ = m.WorkTree()
	}
	status, err := m.Tree.Status()
	if err != nil {
		m.Logger.Do.Errorf("🚨 Couldn't get status:\n\t%v", err)
	}
	return status, err
}
