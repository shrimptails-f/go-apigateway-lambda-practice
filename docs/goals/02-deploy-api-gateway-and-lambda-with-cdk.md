# AWS CDK で API Gateway と Lambda をまとめてデプロイする

network, repository, application を組み立ててスタックに載せる [infra/main.go:43](/home/dev/backend/infra/main.go:43)
```go
network := createNetworkResources(stack)
repository := importEcrRepository(stack)
application := createApplicationResources(stack, network, repository)

awscdk.NewCfnOutput(stack, jsii.String("ApiUrl"), &awscdk.CfnOutputProps{
	Value: application.api.Url(),
})
```

Lambda, authorizer, API Gateway を CDK で作成する [infra/main.go:108](/home/dev/backend/infra/main.go:108)
```go
func createApplicationResources(stack awscdk.Stack, network networkResources, repository awsecr.IRepository) applicationResources {
	fn := awslambda.NewDockerImageFunction(stack, jsii.String("UserFunction"), &awslambda.DockerImageFunctionProps{
		FunctionName: jsii.String("backend-user"),
```
