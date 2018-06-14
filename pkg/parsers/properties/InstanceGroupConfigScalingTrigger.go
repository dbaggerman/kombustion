package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// InstanceGroupConfigScalingTrigger Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingtrigger.html
type InstanceGroupConfigScalingTrigger struct {
	CloudWatchAlarmDefinition *InstanceGroupConfigCloudWatchAlarmDefinition `yaml:"CloudWatchAlarmDefinition"`
}

// InstanceGroupConfigScalingTrigger validation
func (resource InstanceGroupConfigScalingTrigger) Validate() []error {
	errs := []error{}

	if resource.CloudWatchAlarmDefinition == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'CloudWatchAlarmDefinition'"))
	} else {
		errs = append(errs, resource.CloudWatchAlarmDefinition.Validate()...)
	}
	return errs
}