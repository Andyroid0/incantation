package git

func New(path string) Model {
	return Model{
		Path: path,
	}
}
