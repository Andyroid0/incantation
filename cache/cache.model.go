package cache

import (
	"time"
)

type Repos struct {
	Data         []RepoData
	LastModified time.Time
}

type RepoData struct {
	Name string
	Path string
}
