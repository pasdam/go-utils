package jsonschemautils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

func Test_ValidateRaw(t *testing.T) {
	type mocks struct {
		jsonSchemaValidationErr error
	}
	type args struct {
		values     string
		schemaJSON string
	}
	tests := []struct {
		name    string
		mocks   mocks
		args    args
		wantErr string
	}{
		{
			name:  "Should not return error if the validation succeed",
			mocks: mocks{},
			args: args{
				values: `{"name": "John", "age": 30}`,
				schemaJSON: `{
					"$schema": "http://json-schema.org/draft-07/schema#",
					"$id": "http://example.com/person.schema.json",
					"type": "object",
					"properties": {
						"name": {
							"type": "string"
						},
						"age": {
							"type": "integer",
							"minimum": 0
						}
					},
					"required": ["name"]
				}`,
			},
			wantErr: "",
		},
		{
			name:  "Should describe the schema violation",
			mocks: mocks{},
			args: args{
				values: `{"name": "John", "age": -1}`,
				schemaJSON: `{
					"$schema": "http://json-schema.org/draft-07/schema#",
					"$id": "http://example.com/person.schema.json",
					"type": "object",
					"properties": {
						"name": {
							"type": "string"
						},
						"age": {
							"type": "integer",
							"minimum": 0
						}
					},
					"required": ["name"]
				}`,
			},
			wantErr: "- age: Must be greater than or equal to 0\n",
		},
		{
			name: "Should propagate error when json validation throws one",
			mocks: mocks{
				jsonSchemaValidationErr: errors.New("some-json-validation-error"),
			},
			args: args{
				values:     "",
				schemaJSON: "{}",
			},
			wantErr: "some-json-validation-error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mocks.jsonSchemaValidationErr != nil {
				mockValidateJson(t, tt.mocks.jsonSchemaValidationErr)
			}

			err := ValidateRaw([]byte(tt.args.values), []byte(tt.args.schemaJSON))

			if len(tt.wantErr) == 0 {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantErr, err.Error())
			}
		})
	}
}

func mockValidateJson(t *testing.T, err error) {
	originalValue := validateJson
	validateJson = func(ls, ld gojsonschema.JSONLoader) (*gojsonschema.Result, error) {
		return nil, err
	}
	t.Cleanup(func() { validateJson = originalValue })
}
