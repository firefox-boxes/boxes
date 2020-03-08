package boxes

import (
	"errors"
	"os"
	"path/filepath"
)

type ProbeResult struct {
	AppDataDir string
	ProfilesDir string
}

func (pr *ProbeResult) GetRelDir(path string) string {
	return filepath.Join(pr.AppDataDir, path)
}

func (pr *ProbeResult) GetProfileDir(pid ProfileID) string {
	return filepath.Join(pr.ProfilesDir, string(pid))
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
	probeResult.AppDataDir = filepath.Join(homeDir, ".FirefoxBoxes")
	mkdirIfNotExists(probeResult.AppDataDir, os.FileMode(DIR_MODE))
	probeResult.ProfilesDir = filepath.Join(probeResult.AppDataDir, "Boxes")
	mkdirIfNotExists(probeResult.ProfilesDir, os.FileMode(DIR_MODE))
	return probeResult
}