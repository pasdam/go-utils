package filetestutils

import (
	"os"
	"testing"
)

// PathExist is a test helper function used to assert that the specified path
// exists
func PathExist(t *testing.T, filePath string) {
	_, err := os.Stat(filePath)
	if err != nil {
		t.Errorf("Error while accessing expected path %s: %s", filePath, err.Error())
		return
	}
}
