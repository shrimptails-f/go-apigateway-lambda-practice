# Lambda は ECR のコンテナイメージを使って起動する

既存の `backend-user` ECR repository を参照する [infra/main.go:99](/home/dev/backend/infra/main.go:99)
```go
func importEcrRepository(stack awscdk.Stack) awsecr.IRepository {
	return awsecr.Repository_FromRepositoryName(
		stack,
		jsii.String("UserRepository"),
		jsii.String("backend-user"),
	)
}
```

`DockerImageCode_FromEcr` で ECR イメージを Lambda のコードとして使う [infra/main.go:109](/home/dev/backend/infra/main.go:109)
```go
fn := awslambda.NewDockerImageFunction(stack, jsii.String("UserFunction"), &awslambda.DockerImageFunctionProps{
	FunctionName: jsii.String("backend-user"),
	Code: awslambda.DockerImageCode_FromEcr(
		repository,
		&awslambda.EcrImageCodeProps{
			TagOrDigest: jsii.String(ecrV),
		},
	),
```
