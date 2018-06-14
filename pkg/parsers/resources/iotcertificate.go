package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// IoTCertificate Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-certificate.html
type IoTCertificate struct {
	Type       string                   `yaml:"Type"`
	Properties IoTCertificateProperties `yaml:"Properties"`
	Condition  interface{}              `yaml:"Condition,omitempty"`
	Metadata   interface{}              `yaml:"Metadata,omitempty"`
	DependsOn  interface{}              `yaml:"DependsOn,omitempty"`
}

// IoTCertificate Properties
type IoTCertificateProperties struct {
	CertificateSigningRequest interface{} `yaml:"CertificateSigningRequest"`
	Status                    interface{} `yaml:"Status"`
}

// NewIoTCertificate constructor creates a new IoTCertificate
func NewIoTCertificate(properties IoTCertificateProperties, deps ...interface{}) IoTCertificate {
	return IoTCertificate{
		Type:       "AWS::IoT::Certificate",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseIoTCertificate parses IoTCertificate
func ParseIoTCertificate(name string, data string) (cf types.TemplateObject, err error) {
	var resource IoTCertificate
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: IoTCertificate - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseIoTCertificate validator
func (resource IoTCertificate) Validate() []error {
	return resource.Properties.Validate()
}

// ParseIoTCertificateProperties validator
func (resource IoTCertificateProperties) Validate() []error {
	errs := []error{}
	if resource.CertificateSigningRequest == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'CertificateSigningRequest'"))
	}
	if resource.Status == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Status'"))
	}
	return errs
}