package utils

import (
	"fmt"
)

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func FormatFileSize(size int64) string {
	const (
		KB = 1 << (10 * 1) // 1024
		MB = 1 << (10 * 2) // 1048576
		GB = 1 << (10 * 3) // 1073741824
		TB = 1 << (10 * 4) // 1099511627776
	)
	switch {
	case size < KB:
		return fmt.Sprintf("%d b", size)
	case size < MB:
		return fmt.Sprintf("%.1f kb", float64(size)/KB)
	case size < GB:
		return fmt.Sprintf("%.1f mb", float64(size)/MB)
	case size < TB:
		return fmt.Sprintf("%.1f gb", float64(size)/GB)
	default:
		return fmt.Sprintf("%.1f tb", float64(size)/TB)
	}
}
