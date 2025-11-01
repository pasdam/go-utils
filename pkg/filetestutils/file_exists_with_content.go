package filetestutils

import (
	"os"
	"testing"
)

// FileExistsWithContent is a test helper function used to assert that the
// specified file exists and has the expected content
func FileExistsWithContent(t *testing.T, filePath string, expectedContent string) {
	actualContent, err := os.ReadFile(filePath)

	if err != nil {
		t.Errorf("Error while reading expected file %s: %s", filePath, err.Error())
		return
	}

	actualContentString := string(actualContent)
	if expectedContent != actualContentString {
		t.Errorf("Expected content of %s: \"%s\". Got: \"%s\"", filePath, expectedContent, actualContentString)
		return
	}
}
