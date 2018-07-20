package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// SNSTopicPolicy Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-policy.html
type SNSTopicPolicy struct {
	Type       string                   `yaml:"Type"`
	Properties SNSTopicPolicyProperties `yaml:"Properties"`
	Condition  interface{}              `yaml:"Condition,omitempty"`
	Metadata   interface{}              `yaml:"Metadata,omitempty"`
	DependsOn  interface{}              `yaml:"DependsOn,omitempty"`
}

// SNSTopicPolicy Properties
type SNSTopicPolicyProperties struct {
	PolicyDocument interface{} `yaml:"PolicyDocument"`
	Topics         interface{} `yaml:"Topics"`
}

// NewSNSTopicPolicy constructor creates a new SNSTopicPolicy
func NewSNSTopicPolicy(properties SNSTopicPolicyProperties, deps ...interface{}) SNSTopicPolicy {
	return SNSTopicPolicy{
		Type:       "AWS::SNS::TopicPolicy",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseSNSTopicPolicy parses SNSTopicPolicy
func ParseSNSTopicPolicy(
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
	var resource SNSTopicPolicy
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

// ParseSNSTopicPolicy validator
func (resource SNSTopicPolicy) Validate() []error {
	return resource.Properties.Validate()
}

// ParseSNSTopicPolicyProperties validator
func (resource SNSTopicPolicyProperties) Validate() []error {
	errors := []error{}
	if resource.PolicyDocument == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'PolicyDocument'"))
	}
	if resource.Topics == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Topics'"))
	}
	return errors
}