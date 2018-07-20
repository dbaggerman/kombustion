package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// EKSCluster Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-cluster.html
type EKSCluster struct {
	Type       string               `yaml:"Type"`
	Properties EKSClusterProperties `yaml:"Properties"`
	Condition  interface{}          `yaml:"Condition,omitempty"`
	Metadata   interface{}          `yaml:"Metadata,omitempty"`
	DependsOn  interface{}          `yaml:"DependsOn,omitempty"`
}

// EKSCluster Properties
type EKSClusterProperties struct {
	Name               interface{}                           `yaml:"Name,omitempty"`
	RoleArn            interface{}                           `yaml:"RoleArn"`
	Version            interface{}                           `yaml:"Version,omitempty"`
	ResourcesVpcConfig *properties.ClusterResourcesVpcConfig `yaml:"ResourcesVpcConfig"`
}

// NewEKSCluster constructor creates a new EKSCluster
func NewEKSCluster(properties EKSClusterProperties, deps ...interface{}) EKSCluster {
	return EKSCluster{
		Type:       "AWS::EKS::Cluster",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEKSCluster parses EKSCluster
func ParseEKSCluster(
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
	var resource EKSCluster
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

// ParseEKSCluster validator
func (resource EKSCluster) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEKSClusterProperties validator
func (resource EKSClusterProperties) Validate() []error {
	errors := []error{}
	if resource.RoleArn == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'RoleArn'"))
	}
	if resource.ResourcesVpcConfig == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'ResourcesVpcConfig'"))
	} else {
		errors = append(errors, resource.ResourcesVpcConfig.Validate()...)
	}
	return errors
}