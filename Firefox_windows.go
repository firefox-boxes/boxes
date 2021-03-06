// +build windows

package boxes

import (
	"os"
	"path"
	"strings"
)

func gen(n string) string {
	return path.Join(os.Getenv("SYSTEMDRIVE") + "\\Program Files", n, "firefox.exe")
}

func GetPossibilities() []string {
	paths := strings.Split(os.Getenv(firefoxPathEnvName), ";")
	if paths[0] == "" {
		paths = []string{}
	}
	return append([]string{
		gen("Mozilla Firefox"),
		gen("Firefox Developer Edition"),
		gen("Firefox Nightly"),
	}, paths...)
}