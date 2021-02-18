package filenames

import (
	"path"
	"strings"
)

// FileName gets the file name of filepath without the directory or extension
func FileName(fp string) string {
	return strings.TrimSuffix(path.Base(fp), path.Ext(fp))
}
