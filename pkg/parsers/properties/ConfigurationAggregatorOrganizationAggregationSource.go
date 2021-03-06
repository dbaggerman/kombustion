package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ConfigurationAggregatorOrganizationAggregationSource Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationaggregator-organizationaggregationsource.html
type ConfigurationAggregatorOrganizationAggregationSource struct {
	AllAwsRegions interface{} `yaml:"AllAwsRegions,omitempty"`
	RoleArn       interface{} `yaml:"RoleArn"`
	AwsRegions    interface{} `yaml:"AwsRegions,omitempty"`
}

// ConfigurationAggregatorOrganizationAggregationSource validation
func (resource ConfigurationAggregatorOrganizationAggregationSource) Validate() []error {
	errors := []error{}

	if resource.RoleArn == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'RoleArn'"))
	}
	return errors
}
