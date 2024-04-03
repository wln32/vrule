package vrule

import (
	"path/filepath"
	"runtime"
)

func getI18nPath() string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(file), "i18n")
}
