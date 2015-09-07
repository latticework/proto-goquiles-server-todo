package jaliconfig

import (
	"github.com/latticework/proto-goquiles-server-todo/go-jali/jalicore"
	"path/filepath"
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"strings"
)

type JsonConfiguration struct {
	Config JsonConfig
	Json   map[string]interface{}
}


func NewJsonConfiguration(config *JsonConfig) (*JsonConfiguration, error) {
	c := JsonConfiguration{}
	c.Config = buildJsonConfig(config, DefaultJsonConfig)

	settingsFile := filepath.Join(c.Config.Path, c.Config.SettingsFileName)

	fileBytes, err := ioutil.ReadFile(settingsFile)

	if err != nil {
		msg := fmt.Sprintf("Existing settings file '%s' cannont be read: '%s'.", settingsFile, err.Error())
		return _, jalicore.ArgumentError{}.Init("config", msg, err)
	}

	err = json.Unmarshal(fileBytes, &c.Json)

	if err != nil {
		msg := fmt.Sprintf("Settings file '%s' is not a valid JSON document: '%s'.", settingsFile, err.Error())
		return _, jalicore.ArgumentError{}.Init("config", msg, err)
	}

	return &c, _
}

func (config *JsonConfiguration) GetJsonValue(
key string, defaultValue map[string]interface{}) (map[string]interface{}, error) {
	return nil, jalicore.NotImplementedError{}.Init(nil, nil)
}

func (config *JsonConfiguration) GetValue(key string, defaultValue interface{}) (interface{}, error) {
	jsonObject := config.Json
	name := key

	if lastSeparator := strings.LastIndex(key, "."); lastSeparator != -1 {
		baseKey := key[:lastSeparator]
		name = key[lastSeparator:]

		var err error
		jsonObject, err = config.navigateToObject(baseKey)

		if err != nil {
			if err.(*KeyNotFoundError) {
				return defaultValue, _
			}
			return _, err
		}
	}

	var value interface{}
	for _, propertyName := range jsonObject {
		if strings.EqualFold(name, propertyName) {
			value = jsonObject[propertyName]
			break;
		}
	}

	if value != nil {
		return value, _
	}

	if defaultValue != nil {
		return defaultValue, _
	}

	return _, KeyNotFoundError{}.Init(key)
}


func buildJsonConfig(config *JsonConfig, defaultConfig *JsonConfig) (*JsonConfig, error) {
	var newConfig JsonConfig
	if config == nil {
		newConfig = JsonConfig{
			SettingsFileName: defaultConfig.SettingsFileName,
			Path: defaultConfig.Path,
		}
	} else {
		newConfig = JsonConfig{
			SettingsFileName: config.SettingsFileName,
			Path: config.Path,
		}
	}

	if jalicore.IsNilOrEmpty(config.SettingsFileName) {
		newConfig.SettingsFileName = DefaultJsonConfig.SettingsFileName
	}

	if config == nil || newConfig.Path == nil || len(config.Path) == 0 {
		path, err := filepath.Abs(".")

		if err != nil {
			return _, jalicore.InvalidOperationError{}.Init("Could not get the path to the current working directory", _)
		}

		settingsFilePath, err := jalicore.FindSettingsFile(newConfig.SettingsFileName, path)

		if err != nil {
			msg := fmt.Sprintf("Error resolving settings file: '%s'.", err.Error())
			return _, jalicore.InvalidOperationError{}.Init(msg, err)
		}

		if settingsFilePath == nil {
			msg := fmt.Sprintf(
				"Could not find any file named '%s' starting at current directory '%s'.",
				newConfig.SettingsFileName,
				path)

			return _, jalicore.ArgumentError{}.Init("config", msg, _)
		}

		newConfig.Path, newConfig.SettingsFileName = filepath.Split(settingsFilePath)
	} else {
		fileInfo, err := os.Stat(config.Path)

		var msg string

		switch {
		case err == nil && fileInfo.IsDir():
		case err == nil:
			msg = fmt.Sprintf("Path '%s' is not a directory.", newConfig.Path)
		case os.IsNotExist(err):
			msg = fmt.Sprintf("Path '%s' does not exist.", newConfig.Path)
		default:
			msg = fmt.Sprintf("Path '%s' is in error: '%s'.", newConfig.Path, err.Error())
		}

		if msg != nil {
			return _, jalicore.ArgumentError{}.Init("config", msg, err)
		}

		filePath := filepath.Join(config.Path, newConfig.SettingsFileName)

		fileInfo, err = os.Stat(filePath)

		switch {
		case err == nil && !fileInfo.IsDir():
		case err == nil:
			msg = fmt.Sprintf("Settings file '%s' is a directory.", filePath)
		case os.IsNotExist(err):
			msg = fmt.Sprintf("Settings file '%s' does not exist.", filePath)
		default:
			msg = fmt.Sprintf("Settings file '%s' is in error: '%s'.", filePath, err.Error())
		}

		if msg != nil {
			return _, jalicore.ArgumentError{}.Init("config", msg, err)
		}

		absFilePath, err := filepath.Abs(filePath)

		if err != nil {
			msg = fmt.Sprintf("Settings file '%s' is in error: '%s'.", filePath, err.Error())
			return _, jalicore.ArgumentError{}.Init("config", msg, err)
		}

		newConfig.Path, _ = filepath.Split(absFilePath)
	}

	return newConfig
}

func (config *JsonConfiguration) navigateToObject(baseKey string) (map[string]interface{}, error) {
	path := strings.Split(baseKey, ".")

	root := config.Json
	partialPath := []string

	for _, name := range path {
		append(partialPath, name)

		var found bool
		var propertyName string
		for _, propertyName = range root {
			if strings.EqualFold(name, propertyName) {
				found = true
				break;
			}
		}

		if !found {
			return _, KeyNotFoundError{}.Init(strings.Join(partialPath, "."))
		}

		root = root[propertyName]
	}

	return root, _
}