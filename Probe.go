package boxes

import (
	"errors"
	"os"
	"path/filepath"
)

type ProbeResult struct {
	appDataDir string
	profilesDir string
}

func (pr *ProbeResult) GetRelDir(path string) string {
	return filepath.Join(pr.appDataDir, path)
}

func (pr *ProbeResult) GetProfileDir(pid ProfileID) string {
	return filepath.Join(pr.profilesDir, string(pid))
}

func mkdirIfNotExists(path string, fileMode os.FileMode) (string, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, fileMode)
		return path, nil
	}
	return "", errors.New("folder exists")
}

func Probe() ProbeResult {
	probeResult := ProbeResult{}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	probeResult.appDataDir = filepath.Join(homeDir, ".FirefoxBoxes")
	mkdirIfNotExists(probeResult.appDataDir, os.FileMode(DIR_MODE))
	probeResult.profilesDir = filepath.Join(probeResult.appDataDir, "Boxes")
	mkdirIfNotExists(probeResult.profilesDir, os.FileMode(DIR_MODE))
	return probeResult
}