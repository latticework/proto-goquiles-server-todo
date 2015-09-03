package quilesmodel

import (
	"bytes"
	"encoding/json"
	"net/url"
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
