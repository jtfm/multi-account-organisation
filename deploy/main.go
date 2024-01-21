package main

import (
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
	"github.com/jtfm/multi-account-organisation.git/deploy/stacks/root"
	"github.com/jtfm/multi-account-organisation.git/deploy/pkg/core_cdk"
)

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	root.StatefulStack(
		app,
		"RootStatefulStack",
		&root.StatefulStackProps{
			StackProps: awscdk.StackProps{
				Env:         env(),
				Description: jsii.String("A stack containing stateful resources for the root account."),
			},
		})

	root.StatelessStack(
		app,
		"RootStatelessStack",
		&root.StatelessStackProps{
			StackProps: awscdk.StackProps{
				Env:         env(),
				Description: jsii.String("A stack containing stateless resources for the root account."),
			},
		})

	app.Synth(nil)
}
func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
		Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}
