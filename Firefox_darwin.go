// +build darwin

package boxes

func GetPossibilities() [3]string {
	return [3]string{
		"/Applications/Firefox.app/Contents/MacOS/firefox",
		"/Applications/Firefox Nightly.app/Contents/MacOS/firefox",
		"/Applications/Firefox Developer Edition.app/Contents/MacOS/firefox",
	}
}