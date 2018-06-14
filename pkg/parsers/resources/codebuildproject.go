package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// CodeBuildProject Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html
type CodeBuildProject struct {
	Type       string                     `yaml:"Type"`
	Properties CodeBuildProjectProperties `yaml:"Properties"`
	Condition  interface{}                `yaml:"Condition,omitempty"`
	Metadata   interface{}                `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                `yaml:"DependsOn,omitempty"`
}

// CodeBuildProject Properties
type CodeBuildProjectProperties struct {
	BadgeEnabled     interface{}                        `yaml:"BadgeEnabled,omitempty"`
	Description      interface{}                        `yaml:"Description,omitempty"`
	EncryptionKey    interface{}                        `yaml:"EncryptionKey,omitempty"`
	Name             interface{}                        `yaml:"Name,omitempty"`
	ServiceRole      interface{}                        `yaml:"ServiceRole"`
	TimeoutInMinutes interface{}                        `yaml:"TimeoutInMinutes,omitempty"`
	VpcConfig        *properties.ProjectVpcConfig       `yaml:"VpcConfig,omitempty"`
	Source           *properties.ProjectSource          `yaml:"Source"`
	Triggers         *properties.ProjectProjectTriggers `yaml:"Triggers,omitempty"`
	Cache            *properties.ProjectProjectCache    `yaml:"Cache,omitempty"`
	Tags             interface{}                        `yaml:"Tags,omitempty"`
	Environment      *properties.ProjectEnvironment     `yaml:"Environment"`
	Artifacts        *properties.ProjectArtifacts       `yaml:"Artifacts"`
}

// NewCodeBuildProject constructor creates a new CodeBuildProject
func NewCodeBuildProject(properties CodeBuildProjectProperties, deps ...interface{}) CodeBuildProject {
	return CodeBuildProject{
		Type:       "AWS::CodeBuild::Project",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseCodeBuildProject parses CodeBuildProject
func ParseCodeBuildProject(name string, data string) (cf types.TemplateObject, err error) {
	var resource CodeBuildProject
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: CodeBuildProject - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseCodeBuildProject validator
func (resource CodeBuildProject) Validate() []error {
	return resource.Properties.Validate()
}

// ParseCodeBuildProjectProperties validator
func (resource CodeBuildProjectProperties) Validate() []error {
	errs := []error{}
	if resource.ServiceRole == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ServiceRole'"))
	}
	if resource.Source == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Source'"))
	} else {
		errs = append(errs, resource.Source.Validate()...)
	}
	if resource.Environment == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Environment'"))
	} else {
		errs = append(errs, resource.Environment.Validate()...)
	}
	if resource.Artifacts == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Artifacts'"))
	} else {
		errs = append(errs, resource.Artifacts.Validate()...)
	}
	return errs
}