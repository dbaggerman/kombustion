package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ScalingPolicyStepScalingPolicyConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration.html
type ScalingPolicyStepScalingPolicyConfiguration struct {
	AdjustmentType         interface{} `yaml:"AdjustmentType,omitempty"`
	Cooldown               interface{} `yaml:"Cooldown,omitempty"`
	MetricAggregationType  interface{} `yaml:"MetricAggregationType,omitempty"`
	MinAdjustmentMagnitude interface{} `yaml:"MinAdjustmentMagnitude,omitempty"`
	StepAdjustments        interface{} `yaml:"StepAdjustments,omitempty"`
}

// ScalingPolicyStepScalingPolicyConfiguration validation
func (resource ScalingPolicyStepScalingPolicyConfiguration) Validate() []error {
	errors := []error{}

	return errors
}
