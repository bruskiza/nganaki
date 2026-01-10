package utils

import (
	"os"
	"testing"
)


func TestIsGitRepository(t *testing.T) {
	tempDir := t.TempDir()

	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	defer os.Chdir(originalDir)

	// Change to the temporary directory
	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	// Initially, it should not be a git repository
	if IsGitRepository() {
		t.Errorf("Expected not a git repository, but got true")
	}

}

func TestGetLogicalPath(t *testing.T) {
	d := NewDownloader()
	if d.GetLog() != "" {
		t.Errorf("Initial log = %q, want empty", d.GetLog())
	}
}

func TestWriteFileIfNotExists(t *testing.T) {
	tempDir := t.TempDir()
	testFile := tempDir + "/testfile.txt"
	testData := []byte("Hello, World!")

	// Write the file for the first time
	err := WriteFileIfNotExists(testFile, testData)
	if err != nil {
		t.Fatalf("Failed to write file: %v", err)
	}

	// Check if the file exists
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Errorf("Expected file to exist, but it does not")
	}

	// Try writing the file again, it should not overwrite
	err = WriteFileIfNotExists(testFile, []byte("New Data"))
	if err == nil {
		t.Fatalf("This should have errored.")
	}

	// Read the file and check its content
	content, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if string(content) != string(testData) {
		t.Errorf("File content = %q, want %q", string(content), string(testData))
	}
}

