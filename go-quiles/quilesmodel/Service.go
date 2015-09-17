package quilesmodel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/latticework/proto-goquiles-server-todo/go-jali/jalicore"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
)

type Service struct {
	Url       url.URL
	Name      string
	Resources map[string]Resource
}

func (receiver *Service) UnmarshalJSON(data []byte) error {
	var dto struct {
		Url       string   `json:"url"`
		Name      string   `json:"name"`
		Resources []string `json:"resources"`
	}

	decoder := json.NewDecoder(bytes.NewReader(data))

	if err := decoder.Decode(&dto); err != nil {
		return err
	}

	return nil
}

func DecodeService(path string) (*Service, error) {
	var pathAbsolute string

	if path != nil {
		fileInfo, err := os.Stat(path)

		if err != nil {
			msg := fmt.Sprintf("Error accessing path '%s'. Error: ", path, err.Error())
			return _, jalicore.ArgumentError{}.Init(msg, err)
		}

		if fileInfo.IsDir() {
			pathAbsolute, err = filepath.Abs(path)

			if err != nil {
				msg := fmt.Sprintf("Error getting current working directory: '%s'", err.Error())
				return _, jalicore.ArgumentError{}.Init(msg, err)
			}
		}
	} else {
		var err error
		pathAbsolute, err = filepath.Abs("./")

		if err != nil {
			msg := fmt.Sprintf("Error getting current working directory: '%s'", err.Error())
			return _, jalicore.ArgumentError{}.Init(msg, err)
		}
	}

	var serviceFile string

	if pathAbsolute == nil {
		var err error
		serviceFile, err = filepath.Abs(path)

		if err != nil {
			msg := fmt.Sprintf("Error full path of service defintion file: '%s'. Error: '%s'", path, err.Error())
			return _, jalicore.ArgumentError{}.Init(msg, err)
		}
	} else {
		var err error
		serviceFile, err = jalicore.FindSettingsFile("servicequiles.json", pathAbsolute)

		if err != nil {
			var msg string

			switch err := err.(type) {
			case jalicore.StructuredError:
				msg = fmt.Sprintf(
					"Unable to find service definition file from location '%s'.\nError Type: '%s'\nMessage:\n'%s'\nStack:\n%s",
					pathAbsolute, err.TypeName(), err.Error(), err.ErrorStack())
			default:
				msg = fmt.Sprintf(
					"Unable to find service definition file from location '%s'. Error: %s",
					pathAbsolute, err.Error())

			}

			return _, jalicore.ArgumentError{}.Init(msg, err)
		}
	}

	fileBytes, err := ioutil.ReadFile(serviceFile)

	if err != nil {
		msg := fmt.Sprintf("Existing service definition file '%s' cannont be read. Error: '%s'.", serviceFile, err.Error())

		return _, jalicore.ArgumentError{}.Init(msg, err)
	}

	// TODO: Use JSON Schema decoding + error handling
	decoder := json.NewDecoder(bytes.NewReader(fileBytes))

	var service Service

	if err := decoder.Decode(&service); err != nil {
		switch err := err.(type) {
		case ServiceParseError:
			return _, err
		default:
			msg := fmt.Sprintf("Existing service definition file '%s' cannont be parsed. Error: '%s'.", serviceFile, err.Error())
			return _, ServiceParseError{}.Init(msg, err)
		}
	}

	return &service
}
