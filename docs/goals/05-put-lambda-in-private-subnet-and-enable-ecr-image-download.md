# Lambda はプライベートサブネットに配置し、ECR からイメージをダウンロードできること

private isolated subnet を含む VPC を作成する [infra/main.go](/home/dev/backend/infra/main.go)
```go
vpc := awsec2.NewVpc(stack, jsii.String("ApiVpc"), &awsec2.VpcProps{
	SubnetConfiguration: &[]*awsec2.SubnetConfiguration{
		{
			Name:       jsii.String("lambda"),
			SubnetType: awsec2.SubnetType_PRIVATE_ISOLATED,
			CidrMask:   jsii.Number(24),
		},
	},
})
```

ECR API, ECR DKR, S3 の VPC endpoint を作成する [infra/main.go](/home/dev/backend/infra/main.go)
```go
vpc.AddInterfaceEndpoint(jsii.String("EcrApiEndpoint"), &awsec2.InterfaceVpcEndpointOptions{
	Service: awsec2.InterfaceVpcEndpointAwsService_ECR(),
	Subnets: lambdaSubnetSelection,
})

vpc.AddInterfaceEndpoint(jsii.String("EcrDockerEndpoint"), &awsec2.InterfaceVpcEndpointOptions{
	Service: awsec2.InterfaceVpcEndpointAwsService_ECR_DOCKER(),
	Subnets: lambdaSubnetSelection,
})
```

Lambda を VPC の `lambda` subnet group に配置する [infra/main.go](/home/dev/backend/infra/main.go)
```go
fn := awslambda.NewDockerImageFunction(stack, jsii.String("UserFunction"), &awslambda.DockerImageFunctionProps{
	Vpc:        network.vpc,
	VpcSubnets: network.lambdaSubnetSelection,
})
```
