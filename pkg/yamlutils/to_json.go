package yamlutils

import (
	"bytes"

	"sigs.k8s.io/yaml"
)

var yamlToJson = yaml.YAMLToJSON

func ToJSON(yamlData []byte) ([]byte, error) {
	jsonData, err := yamlToJson(yamlData)
	if err != nil {
		return nil, err
	}

	if bytes.Equal(jsonData, []byte("null")) {
		jsonData = []byte("{}")
	}

	return jsonData, nil
}
