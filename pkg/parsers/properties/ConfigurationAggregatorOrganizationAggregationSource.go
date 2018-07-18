package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
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
	errs := []error{}

	if resource.RoleArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoleArn'"))
	}
	return errs
}