package filenames

import (
	"path/filepath"
)

// FileDirectory gets the directory of filepath without the file name
func FileDirectory(fp string) string {
	return filepath.Dir(fp)
}
