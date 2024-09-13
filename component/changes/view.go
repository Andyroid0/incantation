package changes

import (
	// "fmt"
	"strings"

	"github.com/andyroid0/incantation/constants"
	"github.com/andyroid0/incantation/logger"
	"github.com/andyroid0/incantation/shared"

	// "github.com/andyroid0/incantation/utils"
	// "github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

func (m *Model) MenuView() string {
	label := "Commit"
	var btn string
	status, _ := m.Git.Tree.Status()
	if status.IsClean() {
		btn = shared.DisabledButtonStyle().Render(label)
	} else {
		btn = shared.ButtonStyle().Render(label)
	}
	return zone.Mark(constants.Zone_ChangesTab_Button_Commit, btn)
}

func ContentView(logger *logger.Logger) string {
	s := strings.Builder{}

	return s.String()
}
