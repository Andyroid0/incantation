package repos

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) ActionSetPickerList(path string) tea.Cmd {
	dir, err := os.Open(path)
	if err != nil {
		m.Logger.Do.Errorf("🚨 Filepicker error opening file: %v", err)
		return nil
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}

	return func() tea.Msg {
		return MsgSetFilePickerList{
			Value: files,
		}
	}
}

//	func (m Model) ActionSetFilePickerPath(path string) tea.Cmd {
//		return func() tea.Msg {
//			return MsgSetFilePickerList{
//				Value: files,
//			}
//		}
//		m.FilePickerPath
//	}
