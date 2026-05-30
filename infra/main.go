package main

import (
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/jsii-runtime-go"
)

type BackendApiStackProps struct {
	awscdk.StackProps
}

type networkResources struct {
	vpc                   awsec2.IVpc
	lambdaSubnetSelection *awsec2.SubnetSelection
}

type applicationResources struct {
	function awslambda.DockerImageFunction
	api      awsapigateway.LambdaRestApi
}

const ecrV = "1" // 環境変数にするべきだがPoCなので割愛

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	stack := awscdk.NewStack(app, jsii.String("BackendApiStack"), &awscdk.StackProps{
		Env: &awscdk.Environment{
			Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
			Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
		},
	})

	network := createNetworkResources(stack)
	repository := importEcrRepository(stack)
	application := createApplicationResources(stack, network, repository)

	awscdk.NewCfnOutput(stack, jsii.String("ApiUrl"), &awscdk.CfnOutputProps{
		Value: application.api.Url(),
	})

	app.Synth(nil)
}

func createNetworkResources(stack awscdk.Stack) networkResources {
	vpc := awsec2.NewVpc(stack, jsii.String("ApiVpc"), &awsec2.VpcProps{
		MaxAzs:      jsii.Number(2),
		NatGateways: jsii.Number(0),
		SubnetConfiguration: &[]*awsec2.SubnetConfiguration{
			{
				Name:       jsii.String("public"),
				SubnetType: awsec2.SubnetType_PUBLIC,
				CidrMask:   jsii.Number(24),
			},
			{
				Name:       jsii.String("lambda"),
				SubnetType: awsec2.SubnetType_PRIVATE_ISOLATED,
				CidrMask:   jsii.Number(24),
			},
		},
	})

	lambdaSubnetSelection := &awsec2.SubnetSelection{
		SubnetGroupName: jsii.String("lambda"),
	}

	vpc.AddInterfaceEndpoint(jsii.String("EcrApiEndpoint"), &awsec2.InterfaceVpcEndpointOptions{
		Service: awsec2.InterfaceVpcEndpointAwsService_ECR(),
		Subnets: lambdaSubnetSelection,
	})

	vpc.AddInterfaceEndpoint(jsii.String("EcrDockerEndpoint"), &awsec2.InterfaceVpcEndpointOptions{
		Service: awsec2.InterfaceVpcEndpointAwsService_ECR_DOCKER(),
		Subnets: lambdaSubnetSelection,
	})

	vpc.AddGatewayEndpoint(jsii.String("S3Endpoint"), &awsec2.GatewayVpcEndpointOptions{
		Service: awsec2.GatewayVpcEndpointAwsService_S3(),
		Subnets: &[]*awsec2.SubnetSelection{
			lambdaSubnetSelection,
		},
	})

	return networkResources{
		vpc:                   vpc,
		lambdaSubnetSelection: lambdaSubnetSelection,
	}
}

func importEcrRepository(stack awscdk.Stack) awsecr.IRepository {

	return awsecr.Repository_FromRepositoryName(
		stack,
		jsii.String("UserRepository"),
		jsii.String("backend-user"),
	)
}

func createApplicationResources(stack awscdk.Stack, network networkResources, repository awsecr.IRepository) applicationResources {

	fn := awslambda.NewDockerImageFunction(stack, jsii.String("UserFunction"), &awslambda.DockerImageFunctionProps{
		FunctionName: jsii.String("backend-user"),
		Code: awslambda.DockerImageCode_FromEcr(
			repository,
			&awslambda.EcrImageCodeProps{
				TagOrDigest: jsii.String(ecrV),
			},
		),
		Architecture: awslambda.Architecture_X86_64(),
		MemorySize:   jsii.Number(256),
		Timeout:      awscdk.Duration_Seconds(jsii.Number(10)),
		Vpc:          network.vpc,
		VpcSubnets:   network.lambdaSubnetSelection,
	})

	api := awsapigateway.NewLambdaRestApi(stack, jsii.String("UserApi"), &awsapigateway.LambdaRestApiProps{
		Handler: fn,
		Proxy:   jsii.Bool(true),
		DeployOptions: &awsapigateway.StageOptions{
			StageName: jsii.String("prod"),
		},
	})

	return applicationResources{
		function: fn,
		api:      api,
	}
}
