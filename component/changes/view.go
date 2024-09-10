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
	// if disabled {
	// 	return zone.Mark(constants.Zone_ChangesTab_Button_Commit, component.ActiveButtonStyle().Render("Commit"))

	// }
	return zone.Mark(constants.Zone_ChangesTab_Button_Commit, shared.ButtonStyle().Render("Commit"))
}

func ContentView(logger *logger.Logger) string {
	s := strings.Builder{}

	return s.String()
}
