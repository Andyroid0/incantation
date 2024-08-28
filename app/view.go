package app

import (
	"github.com/andyroid0/incantation/component"
	"github.com/andyroid0/incantation/component/repos"
	"github.com/andyroid0/incantation/constants"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

func (m Model) View() string {
	if m.ShowWholePageDialog {
		return zone.Scan(component.WholePageDialog(m.TerminalWidth, m.TerminalHeight, "Select a file."))
	}

	menuWidth, menu :=
		component.MenuPanel(m.TerminalHeight-constants.MenuPanelHeightOffset, constants.Zone_SideBarTab, m.Tabs, m.TabContent, m.ActiveTab)
	contentWidth := m.TerminalWidth - menuWidth - constants.ContentPanelWidthOffset
	contentHeight := m.TerminalHeight - constants.ContentPanelHeightOffset

	var contents string
	if m.ActiveTab == constants.Repos {
		if m.ShowDialog {
			contents = repos.AddDialogView(contentWidth, contentHeight)
		} else {
			m.viewport.Width = contentWidth
			m.viewport.Height = contentHeight
			m.viewport.SetContent(repos.ContentView(m.ReposModel.CursorIndex, m.ReposModel.FilePickerList, m.Logger))
			contents = m.viewport.View()
		}
	} else {
		m.viewport.Width = contentWidth
		m.viewport.Height = contentHeight
		m.viewport.SetContent(
			component.Dialog(contentWidth, contentHeight, m.ShowDialog))
		contents = m.viewport.View()
	}

	return zone.Scan(
		component.
			DocStyle().
			MaxWidth(m.TerminalWidth).
			Render(
				lipgloss.JoinVertical(
					lipgloss.Top,
					component.AppBar(),
					component.Grid(
						menu,
						component.ContentPanel(
							contentWidth,
							contentHeight,
							contents,
						),
					),
					component.StatusBar(m.TerminalWidth-12),
				),
			),
	)
}
