// +build darwin

package boxes

import (
	"os"
	"strings"
)

func GetPossibilities() []string {
	paths := strings.Split(os.Getenv(firefoxPathEnvName), ":")
	if paths[0] == "" {
		paths := []string{}
	}
	return append([]string{
		"/Applications/Firefox.app/Contents/MacOS/firefox",
		"/Applications/Firefox Nightly.app/Contents/MacOS/firefox",
		"/Applications/Firefox Developer Edition.app/Contents/MacOS/firefox"
	}, paths...)
}