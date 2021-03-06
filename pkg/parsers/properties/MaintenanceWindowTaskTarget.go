package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// MaintenanceWindowTaskTarget Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-target.html
type MaintenanceWindowTaskTarget struct {
	Key    interface{} `yaml:"Key"`
	Values interface{} `yaml:"Values,omitempty"`
}

// MaintenanceWindowTaskTarget validation
func (resource MaintenanceWindowTaskTarget) Validate() []error {
	errors := []error{}

	if resource.Key == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Key'"))
	}
	return errors
}
