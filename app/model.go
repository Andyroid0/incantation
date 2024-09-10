package app

import (
	"github.com/andyroid0/incantation/logger"
	"github.com/charmbracelet/bubbles/viewport"

	"github.com/andyroid0/incantation/component/changes"
	"github.com/andyroid0/incantation/component/repos"
	"github.com/andyroid0/incantation/git"
)

type Model struct {
	Logger              *logger.Logger
	ShowWholePageDialog bool
	ShowDialog          bool
	Tabs                []string
	TabContent          []string
	ActiveTab           int
	ReposModel          repos.Model
	ChangesModel        *changes.Model
	viewport            *viewport.Model
	TerminalWidth       int
	TerminalHeight      int
	Git                 *git.Model
}
