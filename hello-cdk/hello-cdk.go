package main

import (
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodecommit"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/pipelines"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type HelloCdkStackProps struct {
	awscdk.StackProps
}

func NewHelloCdkStack(scope constructs.Construct, id string, props *HelloCdkStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		panic("ENVIRONMENT variable is not set")
	}

	awslambda.NewFunction(stack, jsii.String("HelloWorldFunction"), &awslambda.FunctionProps{
		Runtime:    awslambda.Runtime_GO_1_X(),
		Handler:    jsii.String("bootstrap"),
		Code:       awslambda.AssetCode_FromAsset(jsii.String("./lambda"), nil),
		MemorySize: jsii.Number(128),
		Timeout:    awscdk.Duration_Seconds(jsii.Number(30)),
		Environment: &map[string]*string{
			"ENVIRONMENT": jsii.String(env),
		},
		Architecture: awslambda.Architecture_ARM_64(),
		Description:  jsii.String("Hello World Lambda Function in Go"),
		Tracing:      awslambda.Tracing_ACTIVE,
	})

	return stack
}

type PipelineStackProps struct {
	awscdk.StackProps
}

func NewPipelineStack(scope constructs.Construct, id string, props *PipelineStackProps) awscdk.Stack {
	stack := awscdk.NewStack(scope, &id, &props.StackProps)

	// Reference your CodeCommit repository
	repo := awscodecommit.Repository_FromRepositoryName(stack, jsii.String("CodeCommitRepo"), jsii.String("Golang-Lambda-Project"))

	// Define the pipeline
	pipeline := pipelines.NewCodePipeline(stack, jsii.String("Pipeline"), &pipelines.CodePipelineProps{
		PipelineName: jsii.String("HelloCdkPipeline"),
		Synth: pipelines.NewShellStep(jsii.String("SynthStep"), &pipelines.ShellStepProps{
			Input: pipelines.CodePipelineSource_CodeCommit(repo, jsii.String("main"), nil),
			Commands: jsii.Strings(
				"npm install -g aws-cdk",
				"go mod tidy",
				"cdk synth",
			),
		}),
	})

	// Add application stage
	pipeline.AddStage(NewMyApplicationStage(stack, "DeployStage", nil), nil)

	return stack
}

type MyApplicationStage struct {
	awscdk.Stage
}

func NewMyApplicationStage(scope constructs.Construct, id string, props *awscdk.StageProps) awscdk.Stage {
	stage := awscdk.NewStage(scope, &id, props)

	NewHelloCdkStack(stage, "HelloCdkStack", &HelloCdkStackProps{})

	return stage
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewPipelineStack(app, "PipelineStack", &PipelineStackProps{
		StackProps: awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	account := os.Getenv("CDK_DEFAULT_ACCOUNT")
	if account == "" {
		account = os.Getenv("AWS_ACCOUNT_ID")
		if account == "" {
			account = "614768946157"
		}
	}

	region := os.Getenv("CDK_DEFAULT_REGION")
	if region == "" {
		region = os.Getenv("AWS_REGION")
		if region == "" {
			region = "us-east-1"
		}
	}

	return &awscdk.Environment{
		Account: jsii.String(account),
		Region:  jsii.String(region),
	}
}
