package api

import (
	"fmt"
	"testing"

	apiTypes "github.com/KablamoOSS/kombustion/pkg/plugins/api/types"
	"github.com/KablamoOSS/kombustion/types"
	"github.com/stretchr/testify/assert"
	"github.com/vmihailenco/msgpack"
)

func TestRegisterPlugin(t *testing.T) {
	tests := []struct {
		input apiTypes.Config
	}{
		{
			input: apiTypes.Config{
				Name:    "Name",
				Version: "0.1.0",
				Prefix:  "Test",
				Help:    apiTypes.Help{},
			},
		},
	}

	for i, test := range tests {
		assert := assert.New(t)

		testOutput := RegisterPlugin(test.input)

		var testConfig apiTypes.Config
		err := msgpack.Unmarshal(testOutput, &testConfig)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(testConfig, test.input, fmt.Sprintf("Test %d", i))
	}
}
func TestRegisterParser(t *testing.T) {
	tests := []struct {
		input func(
			name string,
			data string,
		) (
			conditions types.TemplateObject,
			metadata types.TemplateObject,
			mappings types.TemplateObject,
			outputs types.TemplateObject,
			parameters types.TemplateObject,
			resources types.TemplateObject,
			transform types.TemplateObject,
			errors []error,
		)
		output apiTypes.PluginParserResult
	}{
		{
			input: func(
				name string,
				data string,
			) (
				conditions types.TemplateObject,
				metadata types.TemplateObject,
				mappings types.TemplateObject,
				outputs types.TemplateObject,
				parameters types.TemplateObject,
				resources types.TemplateObject,
				transform types.TemplateObject,
				errors []error,
			) {
				resources = make(types.TemplateObject)

				resources[name] = types.TemplateObject{
					"Name": name,
					"Data": data,
				}

				return
			},
			output: apiTypes.PluginParserResult{
				Conditions: types.TemplateObject{},
				Metadata:   types.TemplateObject{},
				Mappings:   types.TemplateObject{},
				Outputs:    types.TemplateObject{},
				Parameters: types.TemplateObject{},
				Resources: types.TemplateObject{
					"TestName": map[string]interface{}{
						"Name": "TestName",
						"Data": "TestData",
					},
				},
				Transform: types.TemplateObject{},
				Errors:    []error(nil),
			},
		},
	}

	for i, test := range tests {
		assert := assert.New(t)

		name := "TestName"
		data := "TestData"
		testParserFunc := RegisterParser(test.input)

		testOutput := testParserFunc(name, data)

		var pluginResult apiTypes.PluginParserResult
		err := msgpack.Unmarshal(testOutput, &pluginResult)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(test.output, pluginResult, fmt.Sprintf("Test %d", i))
	}
}
