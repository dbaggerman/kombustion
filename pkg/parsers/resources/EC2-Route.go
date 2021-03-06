package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// EC2Route Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html
type EC2Route struct {
	Type       string             `yaml:"Type"`
	Properties EC2RouteProperties `yaml:"Properties"`
	Condition  interface{}        `yaml:"Condition,omitempty"`
	Metadata   interface{}        `yaml:"Metadata,omitempty"`
	DependsOn  interface{}        `yaml:"DependsOn,omitempty"`
}

// EC2Route Properties
type EC2RouteProperties struct {
	DestinationCidrBlock        interface{} `yaml:"DestinationCidrBlock,omitempty"`
	DestinationIpv6CidrBlock    interface{} `yaml:"DestinationIpv6CidrBlock,omitempty"`
	EgressOnlyInternetGatewayId interface{} `yaml:"EgressOnlyInternetGatewayId,omitempty"`
	GatewayId                   interface{} `yaml:"GatewayId,omitempty"`
	InstanceId                  interface{} `yaml:"InstanceId,omitempty"`
	NatGatewayId                interface{} `yaml:"NatGatewayId,omitempty"`
	NetworkInterfaceId          interface{} `yaml:"NetworkInterfaceId,omitempty"`
	RouteTableId                interface{} `yaml:"RouteTableId"`
	VpcPeeringConnectionId      interface{} `yaml:"VpcPeeringConnectionId,omitempty"`
}

// NewEC2Route constructor creates a new EC2Route
func NewEC2Route(properties EC2RouteProperties, deps ...interface{}) EC2Route {
	return EC2Route{
		Type:       "AWS::EC2::Route",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEC2Route parses EC2Route
func ParseEC2Route(
	name string,
	data string,
) (
	source string,
	conditions types.TemplateObject,
	metadata types.TemplateObject,
	mappings types.TemplateObject,
	outputs types.TemplateObject,
	parameters types.TemplateObject,
	resources types.TemplateObject,
	transform types.TemplateObject,
	errors []error,
) {
	source = "kombustion-core-resources"
	var resource EC2Route
	err := yaml.Unmarshal([]byte(data), &resource)

	if err != nil {
		errors = append(errors, err)
		return
	}

	if validateErrs := resource.Properties.Validate(); len(errors) > 0 {
		errors = append(errors, validateErrs...)
		return
	}

	resources = types.TemplateObject{name: resource}

	return
}

// ParseEC2Route validator
func (resource EC2Route) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEC2RouteProperties validator
func (resource EC2RouteProperties) Validate() []error {
	errors := []error{}
	if resource.RouteTableId == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'RouteTableId'"))
	}
	return errors
}
