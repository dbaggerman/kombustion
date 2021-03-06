package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// DataSourceLambdaConfig Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-lambdaconfig.html
type DataSourceLambdaConfig struct {
	LambdaFunctionArn interface{} `yaml:"LambdaFunctionArn"`
}

// DataSourceLambdaConfig validation
func (resource DataSourceLambdaConfig) Validate() []error {
	errors := []error{}

	if resource.LambdaFunctionArn == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'LambdaFunctionArn'"))
	}
	return errors
}
