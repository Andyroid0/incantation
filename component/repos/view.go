package repos

import (
	"fmt"
	"io/fs"
	"strings"

	"github.com/andyroid0/incantation/constants"
	"github.com/andyroid0/incantation/logger"
	"github.com/andyroid0/incantation/shared"
	"github.com/andyroid0/incantation/utils"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

func MenuView() string {
	return zone.Mark(constants.Zone_ReposTab_Button_Add, shared.ButtonStyle().Render("Add"))
}

func AddDialogView(containerWidth int, containerHeight int) string {
	addExistingButton := shared.ActiveButtonStyle().Render("Add Existing Repo")
	cloneButton := shared.ActiveButtonStyle().Render("Clone Repo")
	createNewButton := shared.ActiveButtonStyle().Render("Create New Repo")

	buttons := lipgloss.JoinVertical(
		lipgloss.Center,
		zone.Mark(constants.Zone_AddExistingRepoButton, addExistingButton),
		"\n",
		zone.Mark(constants.Zone_CloneRepoButton, cloneButton),
		"\n",
		zone.Mark(constants.Zone_CreateNewRepoButton, createNewButton),
	)
	dialog := lipgloss.Place(containerWidth, containerHeight,
		lipgloss.Center, lipgloss.Center,
		shared.DialogBoxStyle().Padding(1, 8).Render(buttons),
		lipgloss.WithWhitespaceChars("𝕩"),
		lipgloss.WithWhitespaceForeground(constants.Subtle),
	)
	return dialog
}

func ContentView(selectionIndex int, files []fs.FileInfo, logger *logger.Logger) string {
	s := strings.Builder{}

	for i, file := range files {
		var selectionIndicatorStyle lipgloss.Style
		var sizeStyle lipgloss.Style
		var nameStyle lipgloss.Style

		if i == selectionIndex {
			selectionIndicatorStyle =
				lipgloss.NewStyle().
					Bold(true).
					Foreground(lipgloss.Color(constants.ColorGreen)).
					Background(lipgloss.Color(constants.ColorSurface0))

			sizeStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(constants.ColorText)).
				Background(lipgloss.Color(constants.ColorSurface0))

			nameStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(constants.ColorPeach)).
				Background(lipgloss.Color(constants.ColorSurface0))
		} else {
			selectionIndicatorStyle =
				lipgloss.NewStyle().
					Bold(true).
					Foreground(lipgloss.Color(constants.ColorGreen))

			sizeStyle =
				lipgloss.NewStyle().
					Foreground(lipgloss.Color(constants.ColorText))

			nameStyle =
				lipgloss.NewStyle().
					Foreground(lipgloss.Color(constants.ColorPeach))
		}

		var fileName string
		if file.IsDir() {
			fileName = "/" + file.Name()
		} else {
			fileName = file.Name()
		}

		if i == selectionIndex {
			str := zone.Mark(fmt.Sprintf("%v%v", constants.Zone_Filepicker_Row, i), selectionIndicatorStyle.Render("> ")+
				sizeStyle.Render(utils.FormatFileSize(file.Size())+" ")+
				nameStyle.Render(fileName)+
				"\n",
			)
			s.WriteString(str)
		} else {

			str := selectionIndicatorStyle.Render("  ") +
				sizeStyle.Render(utils.FormatFileSize(file.Size())+" ") +
				zone.Mark(fmt.Sprintf("%v%v", constants.Zone_Filepicker_Row, i), nameStyle.Render(fileName)) +
				"\n"

			s.WriteString(str)
		}
	}
	return s.String()
}
