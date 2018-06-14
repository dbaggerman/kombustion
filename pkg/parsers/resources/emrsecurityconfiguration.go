package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// EMRSecurityConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-securityconfiguration.html
type EMRSecurityConfiguration struct {
	Type       string                             `yaml:"Type"`
	Properties EMRSecurityConfigurationProperties `yaml:"Properties"`
	Condition  interface{}                        `yaml:"Condition,omitempty"`
	Metadata   interface{}                        `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                        `yaml:"DependsOn,omitempty"`
}

// EMRSecurityConfiguration Properties
type EMRSecurityConfigurationProperties struct {
	Name                  interface{} `yaml:"Name,omitempty"`
	SecurityConfiguration interface{} `yaml:"SecurityConfiguration"`
}

// NewEMRSecurityConfiguration constructor creates a new EMRSecurityConfiguration
func NewEMRSecurityConfiguration(properties EMRSecurityConfigurationProperties, deps ...interface{}) EMRSecurityConfiguration {
	return EMRSecurityConfiguration{
		Type:       "AWS::EMR::SecurityConfiguration",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEMRSecurityConfiguration parses EMRSecurityConfiguration
func ParseEMRSecurityConfiguration(name string, data string) (cf types.TemplateObject, err error) {
	var resource EMRSecurityConfiguration
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EMRSecurityConfiguration - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseEMRSecurityConfiguration validator
func (resource EMRSecurityConfiguration) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEMRSecurityConfigurationProperties validator
func (resource EMRSecurityConfigurationProperties) Validate() []error {
	errs := []error{}
	if resource.SecurityConfiguration == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SecurityConfiguration'"))
	}
	return errs
}