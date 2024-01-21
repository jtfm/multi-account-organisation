package root

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	
)

type StatefulStackProps struct {
	awscdk.StackProps
}

func StatefulStack(
	scope constructs.Construct,
	id string,
	props *StatefulStackProps) awscdk.Stack {

	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	bucket := awss3.NewBucket(
		stack,
		jsii.String("CloudTrailBucket"),
		&awss3.BucketProps{
			RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
			LifecycleRules: &[]*awss3.LifecycleRule{
				{
					Enabled:    jsii.Bool(true),
					Expiration: awscdk.Duration_Days(jsii.Number[float64](90)),
				}}})

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
