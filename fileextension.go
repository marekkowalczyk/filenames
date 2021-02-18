package filenames

import (
	"path"
	"strings"
)

// FileExtension gets the file extension of filepath without the leading dot
func FileExtension(fp string) string {
	return strings.TrimPrefix(path.Ext(fp), ".")
}
