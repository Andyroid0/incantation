package constants

const (
	Changes = iota
	History
	Repos
	Branches
)

func Pages() []string { return []string{"Changes", "History", "Repos", "Branches"} }
