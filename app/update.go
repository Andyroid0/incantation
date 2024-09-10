package app

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/andyroid0/incantation/component/repos"

	"github.com/andyroid0/incantation/constants"
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.TerminalWidth = msg.Width
		m.TerminalHeight = msg.Height
		return m, tea.Batch(cmds...)
	case repos.MsgSetFilePickerList:
		m.ReposModel.FilePickerList = msg.Value
		for i := 0; i < len(m.ReposModel.FilePickerList); i++ {
			if m.ReposModel.FilePickerList[i].Name() == m.ReposModel.PreviousDirName {
				m.ReposModel.CursorIndex = i
				m.viewport.YOffset = i
				break
			}
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "m":
			m.ShowDialog = !m.ShowDialog
		case "down":
			if m.ActiveTab == constants.Repos && m.ReposModel.CursorIndex < len(m.ReposModel.FilePickerList)-1 {
				m.ReposModel.CursorIndex += 1
			}
		case "up":
			if m.ActiveTab == constants.Repos && m.ReposModel.CursorIndex > 0 {
				m.ReposModel.CursorIndex -= 1
			}
		case "enter", "right":
			if m.ActiveTab == constants.Repos {
				file := m.ReposModel.FilePickerList[m.ReposModel.CursorIndex]
				if file.IsDir() {
					m.ReposModel.CursorIndex = 0
					m.viewport.YOffset = 0
					path := m.ReposModel.FilePickerPath + "/" + file.Name()
					m.ReposModel.FilePickerPath = path
					cmds = append(cmds, m.ReposModel.ActionSetPickerList(path))
				}
			}
		case "left":
			if m.ActiveTab == constants.Repos {
				stringArr := strings.Split(m.ReposModel.FilePickerPath, "/")
				length := len(stringArr)
				if length > 2 {
					m.Logger.Do.Print(stringArr)
					m.ReposModel.PreviousDirName = stringArr[length-1]

					stringArr = stringArr[:len(stringArr)-1]
					path := strings.Join(stringArr, "/")
					m.ReposModel.FilePickerPath = path

					cmds = append(cmds, m.ReposModel.ActionSetPickerList(path))
				}
			}
		}

	case tea.MouseMsg:
		if zone.Get(constants.Zone_TestModalOk).InBounds(msg) {
			m.ShowDialog = false
		}
		if zone.Get(constants.Zone_WholePageModalOk).InBounds(msg) {
			dir, err := os.Open(".")
			if err != nil {
				log.Fatal(err)
			}
			defer dir.Close()

			files, err := dir.Readdir(-1)
			if err != nil {
				log.Fatal(err)
			}

			for _, file := range files {
				if file.IsDir() {
					m.Logger.Do.Printf("Directory:\n\tname: %v, size: %v", file.Name(), file.Size())
				} else {
					m.Logger.Do.Printf("File:\n\t%v, size: %v", file.Name(), file.Size())
				}
			}
		}
		if zone.Get(constants.Zone_SideBarTab0).InBounds(msg) {
			m.ActiveTab = constants.Changes
			m.ShowDialog = false
		}
		if zone.Get(constants.Zone_SideBarTab1).InBounds(msg) {
			m.ActiveTab = constants.History
			m.ShowDialog = false
		}
		if zone.Get(constants.Zone_SideBarTab2).InBounds(msg) {
			m.ActiveTab = constants.Repos
			m.ShowDialog = false
		}
		if zone.Get(constants.Zone_SideBarTab3).InBounds(msg) {
			m.ActiveTab = constants.Branches
			m.ShowDialog = false
		}
		if zone.Get(constants.Zone_ReposTab_Button_Add).InBounds(msg) {
			m.ShowDialog = true
		}
		if zone.Get(constants.Zone_AddExistingRepoButton).InBounds(msg) {
			m.SetShowDialog(false)
			cmds = append(
				cmds,
				m.ReposModel.ActionSetPickerList(m.ReposModel.FilePickerPath),
			)
		}
		if zone.Get(constants.Zone_CloneRepoButton).InBounds(msg) {
			m.ShowDialog = false
		}
		if zone.Get(constants.Zone_CreateNewRepoButton).InBounds(msg) {
			m.ShowDialog = false
		}

		for i := 0; i < len(m.ReposModel.FilePickerList); i++ {
			if zone.Get(fmt.Sprintf("%v%v", constants.Zone_Filepicker_Row, i)).InBounds(msg) && msg.Action == 1 {
				m.ReposModel.CursorIndex = i
				file := m.ReposModel.FilePickerList[m.ReposModel.CursorIndex]
				if file.IsDir() {
					m.ReposModel.CursorIndex = 0
					path := m.ReposModel.FilePickerPath + "/" + file.Name()
					m.ReposModel.FilePickerPath = path
					cmds = append(cmds, m.ReposModel.ActionSetPickerList(path))
				}
				break
			}
		}
	}
	m.ChangesModel.Update(msg, &cmds)

	*m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
