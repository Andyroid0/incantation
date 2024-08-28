package repos

import (
	"io/fs"

	"github.com/andyroid0/incantation/logger"
)

type Model struct {
	FilePickerList   []fs.FileInfo
	CursorIndex      int
	PreviousDirName  string
	SelectedFilePath string
	FilePickerPath   string
	Logger           *logger.Logger
}
