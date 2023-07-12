package jsonschemautils

import (
	"encoding/json"
)

func Validate(values map[string]interface{}, schemaJSON []byte) (reterr error) {
	valuesJSON, err := json.Marshal(values)
	if err != nil {
		return err
	}

	return ValidateRaw(valuesJSON, schemaJSON)
}
