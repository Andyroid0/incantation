package repos

import (
	"io/fs"
)

type MsgSetFilePickerList struct {
	Value []fs.FileInfo
}

type MsgSetFilePickerPath struct {
	Value string
}
