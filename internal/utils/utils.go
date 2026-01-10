package utils

import (
	"errors"
	"log/slog"
	"os"
	"path/filepath"
)

// FileExists checks if a file exists and is not a directory before we
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil  {
		slog.Error("error getting file", "filename", filename, "error", err)
		return false
	}
	return true
}

// IsGitRepository checks if the current directory is a git repository
func IsGitRepository() bool {
	path, err := os.Executable()
	if err != nil {
		return false
	}

	exPath := filepath.Dir(path)

	if err != nil {
		return false
	}
	slog.Info("is git repo", "path", path, "exPath", exPath)
	
	return FileExists(exPath + "/.git/")
}

func WriteFileIfNotExists(filename string, data []byte) error {
	if FileExists(filename) {
		slog.Info("file already exists, not overwriting", "filename", filename)
		return errors.New("file already exists")
	}
	return os.WriteFile(filename, data, 0644)
	
}
