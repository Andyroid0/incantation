package component

import (
	"github.com/andyroid0/incantation/constants"
	lipgloss "github.com/charmbracelet/lipgloss"
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
