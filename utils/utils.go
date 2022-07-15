package utils

import (
	"strings"
)

// GetFilenameFromPath
func GetFilenameFromPath(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) == 0 {
		return ""
	}
	return parts[len(parts)-1]
}

// GetDirectoryFromPath
func GetDirectoryFromPath(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) <= 1 {
		return ""
	}
	director := parts[0]
	for i := 1; i < len(parts)-1; i++ {
		director += "/"
		director += parts[i]
	}
	return director
}
