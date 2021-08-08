package utils

import (
	"os"
)

// IsFile Check if it’s a file
func IsFile(name string) bool {
	f, err := os.Stat(name)

	return err == nil && !f.IsDir()
}

// IsDir Check if it’s a directory
func IsDir(name string) bool {
	f, err := os.Stat(name)

	return err == nil && f.IsDir()
}

// IsSymlink Check if it’s a Symlink
func IsSymlink(name string) bool {
	f, err := os.Lstat(name)

	return err == nil && (f.Mode()&os.ModeSymlink != 0)
}

// FileExists Checks whether a file or directory exists
func FileExists(name string) bool {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
