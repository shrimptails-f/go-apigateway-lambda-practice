# AWS CDK で API Gateway と Lambda をまとめてデプロイする

network, repository, application を組み立ててスタックに載せる [infra/main.go](/home/dev/backend/infra/main.go)
```go
network := createNetworkResources(stack)
repository := importEcrRepository(stack)
application := createApplicationResources(stack, network, repository)

awscdk.NewCfnOutput(stack, jsii.String("ApiUrl"), &awscdk.CfnOutputProps{
	Value: application.api.Url(),
})
```

Lambda, authorizer, API Gateway を CDK で作成する [infra/main.go](/home/dev/backend/infra/main.go)
```go
func createApplicationResources(stack awscdk.Stack, network networkResources, repository awsecr.IRepository) applicationResources {
	fn := awslambda.NewDockerImageFunction(stack, jsii.String("UserFunction"), &awslambda.DockerImageFunctionProps{
		FunctionName: jsii.String("backend-user"),
```
