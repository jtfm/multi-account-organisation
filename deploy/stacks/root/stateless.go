package root

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type StatelessStackProps struct {
	awscdk.StackProps
}

func StatelessStack(
	scope constructs.Construct,
	id string,
	props *StatelessStackProps) awscdk.Stack {

	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	bucketArn := awscdk.Fn_ImportValue(jsii.String("CloudTrailBucketArn-CloudTrailBucketArn"))

	outputName := "CloudTrailBucketArn"
	awscdk.NewCfnOutput(
		stack,
		jsii.String(outputName),
		&awscdk.CfnOutputProps{
			Value:      bucket.BucketArn(),
			ExportName: jsii.String(fmt.Sprintf("%s-%s", outputName, *stack.StackName())),
		})

	return stack
}
