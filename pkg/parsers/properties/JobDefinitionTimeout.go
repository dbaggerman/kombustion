package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// JobDefinitionTimeout Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-timeout.html
type JobDefinitionTimeout struct {
	AttemptDurationSeconds interface{} `yaml:"AttemptDurationSeconds,omitempty"`
}

// JobDefinitionTimeout validation
func (resource JobDefinitionTimeout) Validate() []error {
	errors := []error{}

	return errors
}
