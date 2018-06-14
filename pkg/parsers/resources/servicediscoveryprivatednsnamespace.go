package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// ServiceDiscoveryPrivateDnsNamespace Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-privatednsnamespace.html
type ServiceDiscoveryPrivateDnsNamespace struct {
	Type       string                                        `yaml:"Type"`
	Properties ServiceDiscoveryPrivateDnsNamespaceProperties `yaml:"Properties"`
	Condition  interface{}                                   `yaml:"Condition,omitempty"`
	Metadata   interface{}                                   `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                                   `yaml:"DependsOn,omitempty"`
}

// ServiceDiscoveryPrivateDnsNamespace Properties
type ServiceDiscoveryPrivateDnsNamespaceProperties struct {
	Description interface{} `yaml:"Description,omitempty"`
	Name        interface{} `yaml:"Name"`
	Vpc         interface{} `yaml:"Vpc"`
}

// NewServiceDiscoveryPrivateDnsNamespace constructor creates a new ServiceDiscoveryPrivateDnsNamespace
func NewServiceDiscoveryPrivateDnsNamespace(properties ServiceDiscoveryPrivateDnsNamespaceProperties, deps ...interface{}) ServiceDiscoveryPrivateDnsNamespace {
	return ServiceDiscoveryPrivateDnsNamespace{
		Type:       "AWS::ServiceDiscovery::PrivateDnsNamespace",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseServiceDiscoveryPrivateDnsNamespace parses ServiceDiscoveryPrivateDnsNamespace
func ParseServiceDiscoveryPrivateDnsNamespace(name string, data string) (cf types.TemplateObject, err error) {
	var resource ServiceDiscoveryPrivateDnsNamespace
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ServiceDiscoveryPrivateDnsNamespace - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseServiceDiscoveryPrivateDnsNamespace validator
func (resource ServiceDiscoveryPrivateDnsNamespace) Validate() []error {
	return resource.Properties.Validate()
}

// ParseServiceDiscoveryPrivateDnsNamespaceProperties validator
func (resource ServiceDiscoveryPrivateDnsNamespaceProperties) Validate() []error {
	errs := []error{}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.Vpc == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Vpc'"))
	}
	return errs
}