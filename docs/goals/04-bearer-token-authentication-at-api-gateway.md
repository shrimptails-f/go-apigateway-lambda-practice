# Bearer トークン認証を API Gateway の Lambda authorizer で行う

API Gateway の全メソッドに authorizer を適用する [infra/main.go:129](/home/dev/backend/infra/main.go:129)
```go
api := awsapigateway.NewLambdaRestApi(stack, jsii.String("UserApi"), &awsapigateway.LambdaRestApiProps{
	Handler: fn,
	Proxy:   jsii.Bool(true),
	DefaultMethodOptions: &awsapigateway.MethodOptions{
		Authorizer: authorizer,
	},
})
```

Bearer token を検証する Token authorizer 用 Lambda を作成する [infra/main.go:144](/home/dev/backend/infra/main.go:144)
```go
func createAuthorizer(stack awscdk.Stack) awsapigateway.TokenAuthorizer {
	authorizerFunction := awslambda.NewFunction(stack, jsii.String("UserAuthorizerFunction"), &awslambda.FunctionProps{
		Code: awslambda.Code_FromInline(jsii.String(`
POC_TOKENS = {"Poc_tokens"}  # PoC用の固定トークン。実運用では別途シークレット管理を考慮してください。

def handler(event, context):
    auth_header = event.get("authorizationToken", "")
    if not auth_header.startswith("Bearer "):
        raise Exception("Unauthorized")
`)),
```
