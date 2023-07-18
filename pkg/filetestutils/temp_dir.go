package filetestutils

import (
	"io/ioutil"
	"os"
	"testing"
)

// TempDir is a test helper function used to create a temporary dir
func TempDir(t *testing.T) string {
	dir, err := ioutil.TempDir(os.TempDir(), "test-dir")
	if err != nil {
		t.Error("Unable to create temporary directory. ", err)
	}

	t.Cleanup(func() {
		defer os.RemoveAll(dir)
	})

	return dir
}
