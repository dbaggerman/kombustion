package tasks

import (
	"fmt"
	"os"
	"time"

	printer "github.com/KablamoOSS/go-cli-printer"

	"github.com/aws/aws-sdk-go/aws"
	awsCF "github.com/aws/aws-sdk-go/service/cloudformation"
)

// UpsertStack -
func UpsertStack(
	templateBody []byte,
	parameters []*awsCF.Parameter,
	capabilities []*string,
	stackName string,
	cf *awsCF.CloudFormation,
	tags map[string]string,
) {

	var err error
	var action string

	cfTags := make([]*awsCF.Tag, 0)
	for key, value := range tags {
		// Since aws-sdk-go insists on using string pointers, pointers to the
		// loop variables will have their values changed.
		// Creating a copy of the key / value here means we don't end up with
		// all the array elements referencing the same variable (and thus
		// having the same value).
		k := key
		v := value
		cfTags = append(cfTags, &awsCF.Tag{Key: &k, Value: &v})
	}

	// use template from file
	_, err = cf.DescribeStacks(&awsCF.DescribeStacksInput{StackName: aws.String(stackName)})
	if err == nil { //update
		action = "Updating"
		printer.Step(fmt.Sprintf("%s Stack %s:", action, stackName))
		_, err = cf.UpdateStack(&awsCF.UpdateStackInput{
			StackName:    aws.String(stackName),
			TemplateBody: aws.String(string(templateBody)),
			Parameters:   parameters,
			Capabilities: capabilities,
			Tags:         cfTags,
		})
	} else {
		action = "Creating"
		printer.Step(fmt.Sprintf("%s Stack %s:", action, stackName))
		_, err = cf.CreateStack(&awsCF.CreateStackInput{
			StackName:    aws.String(stackName),
			TemplateBody: aws.String(string(templateBody)),
			Parameters:   parameters,
			Capabilities: capabilities,
			Tags:         cfTags,
		})
	}
	checkError(err)

	processUpsert(stackName, action, cf)
}

// UpsertStackViaS3 -
func UpsertStackViaS3(
	templateURL string,
	parameters []*awsCF.Parameter,
	capabilities []*string,
	stackName string,
	cf *awsCF.CloudFormation,
	tags map[string]string,
) {

	var err error
	var action string

	cfTags := make([]*awsCF.Tag, 0)
	for key, value := range tags {
		cfTags = append(cfTags, &awsCF.Tag{Key: &key, Value: &value})
	}

	// use cf template url
	_, err = cf.DescribeStacks(&awsCF.DescribeStacksInput{StackName: aws.String(stackName)})
	if err == nil { //update
		action = "Updating"
		printer.Step(fmt.Sprintf("%s Stack %s:", action, stackName))

		_, err = cf.UpdateStack(&awsCF.UpdateStackInput{
			StackName:    aws.String(stackName),
			TemplateURL:  aws.String(templateURL),
			Parameters:   parameters,
			Capabilities: capabilities,
			Tags:         cfTags,
		})
	} else { //create
		action = "Creating"
		printer.Step(fmt.Sprintf("%s Stack %s:", action, stackName))
		_, err = cf.CreateStack(&awsCF.CreateStackInput{
			StackName:    aws.String(stackName),
			TemplateURL:  aws.String(templateURL),
			Parameters:   parameters,
			Capabilities: capabilities,
			Tags:         cfTags,
		})
	}
	checkError(err)

	processUpsert(stackName, action, cf)

}

func processUpsert(stackName, action string, cf *awsCF.CloudFormation) {
	var err error
	var status *awsCF.DescribeStacksOutput

	PrintStackEventHeader()
	for {
		printer.Progress(action)
		time.Sleep(2 * time.Second)
		status, err = cf.DescribeStacks(&awsCF.DescribeStacksInput{StackName: aws.String(stackName)})
		checkError(err)

		events, err := cf.DescribeStackEvents(&awsCF.DescribeStackEventsInput{StackName: aws.String(stackName)})
		checkError(err)

		if len(status.Stacks) > 0 {

			stack := status.Stacks[0]
			stackStatus := *stack.StackStatus

			if len(events.StackEvents) > 0 {
				PrintStackEvent(events.StackEvents[0], false)
			}

			if stackStatus != awsCF.StackStatusCreateInProgress &&
				stackStatus != awsCF.StackStatusUpdateInProgress &&
				stackStatus != awsCF.StackStatusUpdateCompleteCleanupInProgress {
				if stackStatus == awsCF.StackStatusCreateComplete ||
					stackStatus == awsCF.StackStatusUpdateComplete {
					printer.SubStep(
						fmt.Sprintf("Success %s Stack %s", action, stackName),
						1,
						true,
						true,
					)
					os.Exit(0)
				} else {
					printer.SubStep(
						fmt.Sprintf("Fail %s Stack %s", action, stackName),
						1,
						true,
						true,
					)

					printer.Error(fmt.Errorf("Upsert Failed"), "", "")
					time.Sleep(2 * time.Second)
					os.Exit(1)
				}
			}
		}
	}
}
