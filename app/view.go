package app

import (
	"github.com/andyroid0/incantation/component/repos"
	"github.com/andyroid0/incantation/shared"

	"github.com/andyroid0/incantation/constants"

	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

func (m Model) View() string {
	if m.ShowWholePageDialog {
		return zone.Scan(shared.WholePageDialog(m.TerminalWidth, m.TerminalHeight, "Select a file."))
	}

	menuWidth, menu :=
		shared.MenuPanel(
			m.TerminalHeight-constants.MenuPanelHeightOffset,
			constants.Zone_SideBarTab,
			m.Tabs,
			m.TabContent,
			m.ActiveTab,
		)
	contentWidth := m.TerminalWidth - menuWidth - constants.ContentPanelWidthOffset
	contentHeight := m.TerminalHeight - constants.ContentPanelHeightOffset

	var contents string
	if m.ActiveTab == constants.Repos {
		if m.ShowDialog {
			contents = repos.AddDialogView(contentWidth, contentHeight)
		} else {
			m.viewport.Width = contentWidth
			m.viewport.Height = contentHeight - 1 //subtract one line for a top bar to display bread crumbs (path) and a back button
			m.viewport.SetContent(
				repos.ContentView(m.ReposModel.CursorIndex, m.ReposModel.FilePickerList, m.Logger),
			)
			contents = "⬅ Back top bar test. + Add Repository\n" + m.viewport.View()
		}
	} else {
		m.viewport.Width = contentWidth
		m.viewport.Height = contentHeight
		m.viewport.SetContent(
			shared.Dialog(contentWidth, contentHeight, m.ShowDialog))
		contents = m.viewport.View()
	}

	return zone.Scan(
		shared.
			DocStyle().
			MaxWidth(m.TerminalWidth).
			Render(
				lipgloss.JoinVertical(
					lipgloss.Top,
					shared.AppBar(),
					shared.Grid(
						menu,
						shared.ContentPanel(
							contentWidth,
							contentHeight,
							contents,
						),
					),
					shared.StatusBar(m.TerminalWidth-12),
				),
			),
	)
}
