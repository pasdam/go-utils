package filetestutils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TempFileWithContent is a test helper function used to create a temporary file
// with a specified content
func TempFileWithContent(t *testing.T, name string, content []byte) string {
	path := TempFile(t, name)

	err := os.WriteFile(path, content, 0644)
	assert.NoError(t, err)

	return path
}
