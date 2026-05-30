# Bearer トークン認証を API Gateway の Lambda authorizer で行う

API Gateway の全メソッドに authorizer を適用する [infra/main.go](/home/dev/backend/infra/main.go)
```go
api := awsapigateway.NewLambdaRestApi(stack, jsii.String("UserApi"), &awsapigateway.LambdaRestApiProps{
	Handler: fn,
	Proxy:   jsii.Bool(true),
	DefaultMethodOptions: &awsapigateway.MethodOptions{
		Authorizer: authorizer,
	},
})
```

Bearer token を検証する Token authorizer 用 Lambda を作成する [infra/main.go](/home/dev/backend/infra/main.go)
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

## 技術選定理由

- 今回の API は `awsapigateway.LambdaRestApi` を使った REST API なので、そのまま Bearer token を入口で検証するには Lambda authorizer が最も素直。
- 認証失敗を backend-user Lambda まで流さずに API Gateway で遮断できるので、業務処理と認証の責務を分けやすい。
- 実務でも固定トークン比較で動かし、将来は authorizer 内の比較処理だけを Secrets Manager 参照へ差し替えやすい。
