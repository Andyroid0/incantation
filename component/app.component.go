package component

import (
	"fmt"
	"strings"

	"github.com/andyroid0/incantation/constants"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}

var (
	inactiveTabBorder = tabBorderWithBottom("┴", "─", "┴")
	activeTabBorder   = tabBorderWithBottom("┘", " ", "└")
	docStyle          = lipgloss.NewStyle().Padding(1, 2, 0, 0)
	inactiveTabStyle  = lipgloss.NewStyle().Border(inactiveTabBorder, true).BorderForeground(constants.Hightlight).Padding(0, 1)
	activeTabStyle    = inactiveTabStyle.Border(activeTabBorder, true)
	windowStyle       = lipgloss.NewStyle().
				BorderForeground(constants.Hightlight).
				Padding(2, 0).Align(lipgloss.Center).
				Border(lipgloss.NormalBorder()).
				UnsetBorderTop()
)

func AppBar() string {
	return lipgloss.JoinHorizontal(lipgloss.Left,
		zone.Mark(constants.Zone_AppBar_Button_Fetch, ActiveButtonStyle().Margin(0, 1).Render("Fetch Origin")),
		zone.Mark(constants.Zone_AppBar_Button_Push, ActiveButtonStyle().Margin(0, 1).Render("Push")),
		zone.Mark(constants.Zone_AppBar_Button_Pull, ActiveButtonStyle().Margin(0, 1).Render("Pull")),
		zone.Mark(constants.Zone_AppBar_Button_Stash, ActiveButtonStyle().Margin(0, 1).Render("Stash")),
	)
}

func ContentPanel(width int, height int, content string) string {
	return docStyle.Render(lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(constants.Hightlight).
		Width(width).
		Height(height).Render(content))
}

func Dialog(containerWidth int, containerHeight int, showDialog bool) string {
	if showDialog {
		okButton := ActiveButtonStyle().Render("Yes")
		cancelButton := ButtonStyle().Render("Maybe")

		question := lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render("Are you sure you want to eat marmalade?")
		buttons := lipgloss.JoinHorizontal(lipgloss.Top, zone.Mark(constants.Zone_TestModalOk, okButton), cancelButton)
		ui := lipgloss.JoinVertical(lipgloss.Center, question, buttons)

		dialog := lipgloss.Place(containerWidth, containerHeight,
			lipgloss.Center, lipgloss.Center,
			DialogBoxStyle().Render(ui),
			lipgloss.WithWhitespaceChars("𝕩"),
			lipgloss.WithWhitespaceForeground(Subtle),
		)
		return dialog
	} else {
		return ""
	}
}

func WholePageDialog(containerWidth int, containerHeight int, content string) string {
	okButton := ActiveButtonStyle().Render("Yes")
	cancelButton := ButtonStyle().Render("Maybe")

	question := lipgloss.NewStyle().
		Width(containerWidth / 3 * 2).
		MaxHeight(containerHeight / 3 * 2).
		Align(lipgloss.Center).
		Render(content)
	buttons := lipgloss.JoinHorizontal(lipgloss.Top, zone.Mark(constants.Zone_WholePageModalOk, okButton), cancelButton)
	ui := lipgloss.JoinVertical(lipgloss.Center, question, buttons)

	dialog := lipgloss.Place(containerWidth, containerHeight,
		lipgloss.Center, lipgloss.Center,
		DialogBoxStyle().Render(ui),
		lipgloss.WithWhitespaceChars("𝕩"),
		lipgloss.WithWhitespaceForeground(Subtle),
	)
	return dialog
}

func Grid(menu string, headerAndContentPanel string) string {
	return lipgloss.
		JoinHorizontal(
			lipgloss.Top,
			menu,
			headerAndContentPanel)
}

func MenuPanel(height int, id string, tabs []string, tabContent []string, activeTab int) (int, string) {
	doc := strings.Builder{}

	var renderedTabs []string

	for i, t := range tabs {
		var style lipgloss.Style
		isFirst, isLast, isActive := i == 0, i == len(tabs)-1, i == activeTab
		if isActive {
			style = activeTabStyle
		} else {
			style = inactiveTabStyle
		}
		border, _, _, _, _ := style.GetBorder()
		if isFirst && isActive {
			border.BottomLeft = "│"
		} else if isFirst && !isActive {
			border.BottomLeft = "├"
		} else if isLast && isActive {
			border.BottomRight = "│"
		} else if isLast && !isActive {
			border.BottomRight = "┤"
		}
		style = style.Border(border)
		var label string
		if isActive {
			label = lipgloss.NewStyle().Foreground(lipgloss.Color(constants.ColorMauve)).Bold(true).Render(t)
		} else {
			label = lipgloss.NewStyle().Foreground(lipgloss.Color(constants.ColorText)).Render(t)
		}
		renderedTabs = append(renderedTabs, zone.Mark(id+fmt.Sprintf("%v", i), style.Render(label)))
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
	doc.WriteString(row)
	doc.WriteString("\n")
	doc.WriteString(windowStyle.Height(height).Width(
		(lipgloss.Width(row) - windowStyle.GetHorizontalFrameSize()),
	).Render(tabContent[activeTab]))
	result := docStyle.Render(doc.String())
	resultWidth := lipgloss.Width(result)
	return resultWidth, result
}

func StatusBar(width int) string {
	statusKey := StatusStyle().Render("STATUS")
	encoding := EncodingStyle().Render("UTF-8")
	fishCake := FishCakeStyle().Render("✨ Incantation")
	statusVal := StatusText().
		Width(width - lipgloss.Width(statusKey) - lipgloss.Width(encoding) - lipgloss.Width(fishCake)).
		Render("Ravishing")

	bar := lipgloss.JoinHorizontal(lipgloss.Top,
		statusKey,
		statusVal,
		encoding,
		fishCake,
	)

	return StatusBarStyle().Width(width).Margin(0, 1).Render(bar)
}
