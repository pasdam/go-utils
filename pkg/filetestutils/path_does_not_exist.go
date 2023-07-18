package filetestutils

import (
	"os"
	"testing"
)

// PathDoesNotExist is a test helper function used to assert that the specified
// path does not exist
func PathDoesNotExist(t *testing.T, filePath string) {
	_, err := os.Stat(filePath)
	if err == nil {
		t.Errorf("File %s exists", filePath)
		return
	}
}
