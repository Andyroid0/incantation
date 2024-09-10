package app

import (
	"log"
	"os"

	"github.com/andyroid0/incantation/component/changes"
	"github.com/andyroid0/incantation/component/repos"
	"github.com/andyroid0/incantation/constants"
	"github.com/andyroid0/incantation/git"
	"github.com/andyroid0/incantation/logger"
	"github.com/charmbracelet/bubbles/viewport"

	"golang.org/x/term"
)

func New(logger *logger.Logger) Model {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("🚨 Error, unable to retrieve users home directory:\n\t%v", err)
	}

	physicalWidth, physicalHeight, sizeErr := term.GetSize(int(os.Stdout.Fd()))
	if sizeErr != nil {
		log.Fatalf("🚨 Error, unable to get terminal window dimensions:\n\t%v", err)
	}

	vp := viewport.New(0, physicalHeight-constants.ContentPanelHeightOffset)
	gt := git.New(".", logger)
	model := Model{
		Logger:              logger,
		Tabs:                constants.Pages(),
		ShowWholePageDialog: false,
		ShowDialog:          false,
		ActiveTab:           constants.Changes,
		ReposModel: repos.Model{
			CursorIndex:    0,
			FilePickerPath: homeDir,
		},
		ChangesModel: &changes.Model{
			Logger: logger,
			Git:    &gt,
		},
		TerminalWidth:  physicalWidth,
		TerminalHeight: physicalHeight,
		viewport:       &vp,
		Git:            &gt,
	}
	model.TabContent = []string{
		model.ChangesModel.MenuView(),
		"History Tab",
		repos.MenuView(),
		"Branches Tab"}
	return model
}
