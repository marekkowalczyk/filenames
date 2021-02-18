package filenames

import (
	"path/filepath"
)

// FileBase gets the file name with extension of filepath but without the directory
func FileBase(fp string) string {
	return filepath.Base(fp)
}
