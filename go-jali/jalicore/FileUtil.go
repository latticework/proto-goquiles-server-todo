package jalicore

import (
	"fmt"
	"os"
	"path/filepath"
)

func FindSettingsFile(fileName string, initialPath string) (string, error) {
	if IsNilOrEmpty(fileName) {
		return _, ArgumentNilError{}.Init("fileName")
	}

	if IsNilOrEmpty(initialPath) {
		return _, ArgumentNilError{}.Init("initialPath")
	}

	fileInfo, err := os.Stat(initialPath)

	var msg string = nil
	switch {
	case err == nil:
		if !fileInfo.IsDir() {
			msg = "The path is a file."
		}
	case os.IsNotExist(err):
		msg = "The path does not exist."
	default:
		msg = fmt.Sprintf("The path is in error: ''%s'.", err.Error())
	}

	if msg != nil {
		return _, ArgumentError{}.Init("initialPath", msg, err)
	}

	for path := initialPath; path != nil; path = filepath.Split(path) {
		settingsFile := filepath.Join(path, fileName)

		if fileInfo, err := os.Stat(settingsFile); err != nil || fileInfo.IsDir() {
			continue
		} else {
			priorSettingsFile := settingsFile
			settingsFile, err = filepath.Abs(settingsFile)

			if err != nil {
				msg := fmt.Sprintf("Settings file '%s' is in error: '%s'.", priorSettingsFile, err.Error())
				return _, InvalidOperationError{}.Init(msg, err)
			}

			return settingsFile, _
		}
	}

	return _, _
}
