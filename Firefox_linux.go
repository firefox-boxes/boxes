// +build linux

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
		"/usr/bin/firefox"
	}, paths...)
}