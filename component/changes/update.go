package changes

import (
	// "time"

	"github.com/andyroid0/incantation/constants"
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
	// "github.com/go-git/go-git/v5"
	// "github.com/go-git/go-git/v5/plumbing/object"
)

func (m *Model) Update(msg tea.Msg, cmds *[]tea.Cmd) {
	switch msg := msg.(type) {
	case tea.MouseMsg:
		if msg.Action == tea.MouseActionRelease {
			if zone.Get(constants.Zone_ChangesTab_Button_Commit).InBounds(msg) {
				status, _ := m.Git.GetStatus()
				if status.IsClean() {
					// bails out if there isn't anything to commit.
					return
				}

				// _, additionErr := tree.Add(".")
				// if additionErr != nil {
				// 	m.Logger.Do.Errorf("🚨 Couldn't add changes:\n\t%v", additionErr)
				// }

				// m.Git.Tree.Add(".")

				// commit, commitErr := tree.Commit("Testing git committing from incantation.", &git.CommitOptions{
				// 	Author: &object.Signature{
				// 		Name:  "Andrew Valley",
				// 		Email: "andrewpvalley@gmail.com",
				// 		When:  time.Now(),
				// 	},
				// })
				// if commitErr != nil {
				// 	m.Logger.Do.Errorf("🚨 Couldn't commit changes:\n\t%v", commitErr)
				// }

				// m.Logger.Do.Printf("Commit hash: %s", commit.String())
			}
		}

	}
}
