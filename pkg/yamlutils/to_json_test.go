package yamlutils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ToJSON(t *testing.T) {
	type mocks struct {
		yamlToJsonErr error
	}
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		mocks   mocks
		args    args
		want    string
		wantErr string
	}{
		{
			name:  "Should not return error if the validation succeed",
			mocks: mocks{},
			args: args{
				data: "name: John\nage: 30",
			},
			want:    "{\"age\":30,\"name\":\"John\"}",
			wantErr: "",
		},
		{
			name:  "Should return valid JSON if input is empty",
			mocks: mocks{},
			args: args{
				data: "",
			},
			want:    "{}",
			wantErr: "",
		},
		{
			name: "Should propagate error when YAML to JSON conversion throws it",
			mocks: mocks{
				yamlToJsonErr: errors.New("some-yaml-to-json-error"),
			},
			args: args{
				data: "",
			},
			wantErr: "some-yaml-to-json-error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mocks.yamlToJsonErr != nil {
				mockYamlToJson(t, tt.mocks.yamlToJsonErr)
			}

			jsonData, err := ToJSON([]byte(tt.args.data))

			assert.Equal(t, tt.want, string(jsonData))
			if len(tt.wantErr) == 0 {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantErr, err.Error())
			}
		})
	}
}

func mockYamlToJson(t *testing.T, err error) {
	originalValue := yamlToJson
	yamlToJson = func(y []byte) ([]byte, error) {
		return nil, err
	}
	t.Cleanup(func() { yamlToJson = originalValue })
}
