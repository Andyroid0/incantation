package git

import (
	"errors"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-git/go-git/v5"
)

func (m *Model) Update(msg tea.Msg, cmds *[]tea.Cmd) {
	switch msg := msg.(type) {
	case MsgSetRepo:
		m.Path = msg.path

		r, err := git.PlainOpen(msg.path)
		if errors.Is(err, git.ErrRepositoryNotExists) {
			m.Logger.Do.Errorf("🚨 Error, couldn't open repo. It doesn't exist:\n\t%v", err)
		} else if err != nil {
			m.Logger.Do.Errorf("🚨 Error, couldn't open repo:\n\t%v", err)
		}
		t, err := r.Worktree()
		if err != nil {
			m.Logger.Do.Errorf("🚨 Couldn't get Worktree:\n\t%v", err)
		}

		m.Tree = t
		m.Repository = r

	}
}
