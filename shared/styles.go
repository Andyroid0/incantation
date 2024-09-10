package shared

import (
	"fmt"
	"strings"

	"github.com/andyroid0/incantation/constants"
	lipgloss "github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

// General.
var (
	Subtle    = lipgloss.AdaptiveColor{Light: constants.ColorBase, Dark: constants.ColorBase}
	Highlight = lipgloss.AdaptiveColor{Light: constants.ColorSurface2, Dark: constants.ColorSurface2}
	Special   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}

	ActiveTabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      " ",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┘",
		BottomRight: "└",
	}

	TabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┴",
		BottomRight: "┴",
	}

	PaperBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┘",
		BottomRight: "└",
	}
)

func Divider() string {
	return lipgloss.NewStyle().
		SetString("•").
		Padding(0, 1).
		Foreground(Subtle).
		String()
}

func Url() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(Special)
}

// Tabs.

func Tab() lipgloss.Style {
	return lipgloss.NewStyle().
		Border(TabBorder, true).
		BorderForeground(Highlight).
		Padding(0, 1)
}

func ActiveTab() lipgloss.Style {
	return Tab().Border(ActiveTabBorder, true)
}

func TabGap() lipgloss.Style {
	return Tab().
		BorderTop(false).
		BorderLeft(false).
		BorderRight(false)
}

// Title.

func TitleStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		MarginLeft(1).
		MarginRight(5).
		Padding(0, 1).
		Italic(true).
		Foreground(lipgloss.Color("#FFF7DB")).
		SetString("Lip Gloss")
}

func DescStyle() lipgloss.Style {
	return lipgloss.NewStyle().MarginTop(1)
}

func InfoStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderTop(true).
		BorderForeground(Subtle)
}

func DialogBoxStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(constants.ColorMauve)).
		Padding(1, 0).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true)
}

func ButtonStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(constants.ColorBase)).
		Background(lipgloss.Color(constants.ColorFlamingo)).
		Bold(true).
		Padding(0, 3).
		MarginTop(1)
}

func ActiveButtonStyle() lipgloss.Style {
	return ButtonStyle().
		Foreground(lipgloss.Color(constants.ColorBase)).
		Background(lipgloss.Color(constants.ColorMauve)).
		MarginRight(2).
		Underline(true)
}

func DisabledButtonStyle() lipgloss.Style {
	return ButtonStyle().
		Foreground(lipgloss.Color(constants.ColorBase)).
		Background(lipgloss.Color(constants.ColorSubtext0)).
		MarginRight(2).
		Underline(true)
}

// List.

func List(columnWidth int) lipgloss.Style {
	return lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, true, false, false).
		BorderForeground(Subtle).
		MarginRight(2).
		Height(8).
		Width(columnWidth + 1)
}

func ListHeader() lipgloss.Style {
	return lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		BorderForeground(Subtle).
		MarginRight(2)
}

func ListItem() lipgloss.Style {
	return lipgloss.NewStyle().PaddingLeft(2)
}

func CheckMark() string {
	return lipgloss.NewStyle().SetString("✓").
		Foreground(Special).
		PaddingRight(1).
		String()
}

func ListDone(s string) string {
	return CheckMark() + lipgloss.NewStyle().
		Strikethrough(true).
		Foreground(lipgloss.AdaptiveColor{Light: "#969B86", Dark: "#696969"}).
		Render(s)
}

// Paragraphs/History.

func HistoryStyle(columnWidth int) lipgloss.Style {
	return lipgloss.NewStyle().
		Align(lipgloss.Left).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(Highlight).
		Margin(1, 3, 0, 0).
		Padding(1, 2).
		Height(19).
		Width(columnWidth)
}

// Status Bar.

func StatusNugget() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFDF5")).
		Padding(0, 1)
}

func StatusBarStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
		Background(lipgloss.AdaptiveColor{Light: constants.ColorBase, Dark: constants.ColorBase})
}

func StatusStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Inherit(StatusBarStyle()).
		Foreground(lipgloss.Color(constants.ColorSubtext0)).
		Background(lipgloss.Color(constants.ColorSurface2)).
		Padding(0, 1).
		MarginRight(1)
}

func EncodingStyle() lipgloss.Style {
	return StatusNugget().
		Background(lipgloss.Color(constants.ColorSurface0)).
		Foreground(lipgloss.Color(constants.ColorSubtext0)).
		Align(lipgloss.Right)
}

func StatusText() lipgloss.Style {
	return lipgloss.NewStyle().Inherit(StatusBarStyle())
}

func FishCakeStyle() lipgloss.Style {
	return StatusNugget().Background(lipgloss.Color(constants.ColorBase))
}

// Page.

func DocStyle() lipgloss.Style {
	return lipgloss.NewStyle().Padding(0, 2, 1, 2)
}

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
