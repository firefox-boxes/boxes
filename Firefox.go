package boxes

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const firefoxPathEnvName = "BOXES_FIREFOX_PATH"

const TIMES_JSON = `{
"created": {msTimestamp},
"firstUse": null
}
`

type ProfileID string

type Installation struct {
	Exec string
}

func (I Installation) IExists() bool {
	return exists(I.Exec)
}

func genTimestamp() int64 {
	return time.Now().UnixNano() / int64(1e+6)
}

func CreateProfile(path string) error {
	if exists(path) {
		return errors.New("profile already exists")
	}
	os.Mkdir(path, os.FileMode(DIR_MODE))
	content := strings.NewReplacer("{msTimestamp}", strconv.FormatInt(genTimestamp(), 10)).Replace(TIMES_JSON)
	f, err := os.OpenFile(filepath.Join(path, "times.json"), os.O_WRONLY|os.O_CREATE, os.FileMode(LOCK_FILE_MODE))
	if err != nil {
		return err
	}
	f.WriteString(content)
	f.Close()
	pl, err := os.OpenFile(filepath.Join(path, ".parentlock"), os.O_WRONLY|os.O_CREATE, os.FileMode(FILE_MODE))
	if err != nil {
		return err
	}
	pl.Close()
	return nil
}

func exists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		return false
	}
}

func GetInstallations() []Installation {
	installations := make([]Installation, 0, 3)
	possibilities := GetPossibilities()
	for i := 0; i < len(possibilities); i++ {
		if exists(possibilities[i]) {
			installations = append(installations, Installation{Exec:possibilities[i]})
		}
	}
	return installations
}