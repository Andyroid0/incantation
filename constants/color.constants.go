package constants

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	ColorRosewater = "#f2d5cf"
	ColorFlamingo  = "#eebebe"
	ColorPink      = "#f4b8e4"
	ColorMauve     = "#ca9ee6"
	ColorRed       = "#e78284"
	ColorMaroon    = "#ea999c"
	ColorPeach     = "#ef9f76"
	ColorYellow    = "#e5c890"
	ColorGreen     = "#a6d189"
	ColorTeal      = "#81c8be"
	ColorSky       = "#99d1db"
	ColorSapphire  = "#85c1dc"
	ColorBlue      = "#8caaee"
	ColorLavender  = "#babbf1"
	ColorText      = "#c6d0f5"
	ColorSubtext1  = "#b5bfe2"
	ColorSubtext0  = "#a5adce"
	ColorOverlay2  = "#949cbb"
	ColorOverlay1  = "#838ba7"
	ColorOverlay0  = "#737994"
	ColorSurface2  = "#626880"
	ColorSurface1  = "#51576d"
	ColorSurface0  = "#414559"
	ColorBase      = "#303446"
	ColorMantle    = "#292c3c"
	ColorCrust     = "#232634"
)

var (
	Subtle     = lipgloss.AdaptiveColor{Light: ColorBase, Dark: ColorBase}
	Hightlight = lipgloss.AdaptiveColor{Light: ColorSky, Dark: ColorSky}
)
