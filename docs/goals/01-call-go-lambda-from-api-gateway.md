# Go で書いた Lambda を API Gateway から呼び出す

API Gateway と Lambda をデプロイする [infra/main.go](/home/dev/backend/infra/main.go)
```go
api := awsapigateway.NewLambdaRestApi(stack, jsii.String("UserApi"), &awsapigateway.LambdaRestApiProps{
	Handler: fn,
	Proxy:   jsii.Bool(true),
	DefaultMethodOptions: &awsapigateway.MethodOptions{
		Authorizer: authorizer,
	},
})
```

Lambda handler を起動エントリポイントで登録する [cmd/api/user/main.go](/home/dev/backend/cmd/api/user/main.go)
```go
func main() {
	repo := repository.NewInMemoryUserRepository()
	handler := user.NewHandler(
		application.NewListUseCase(repo),
		application.NewGetDetailUseCase(repo),
	)

	lambda.Start(handler.Handle)
}
```

`/users` と `/user/{id}` の path を解析する [internal/presentation/user/router.go](/home/dev/backend/internal/presentation/user/router.go)
```go
func parseRoute(req events.APIGatewayProxyRequest) (route, error) {
	path := strings.Trim(req.Path, "/")
	if path == "users" {
		return route{kind: routeListUsers}, nil
	}

	parts := strings.Split(path, "/")
	if len(parts) == 2 && parts[0] == "user" && parts[1] != "" {
		return route{kind: routeGetUserDetail, userID: parts[1]}, nil
	}
```

APIGatewayProxyRequestの内容はこうなっている。
```go
// APIGatewayProxyRequest contains data coming from the API Gateway proxy
type APIGatewayProxyRequest struct {
	Resource                        string                        `json:"resource"` // The resource path defined in API Gateway
	Path                            string                        `json:"path"`     // The url path for the caller
	HTTPMethod                      string                        `json:"httpMethod"`
	Headers                         map[string]string             `json:"headers"`
	MultiValueHeaders               map[string][]string           `json:"multiValueHeaders"`
	QueryStringParameters           map[string]string             `json:"queryStringParameters"`
	MultiValueQueryStringParameters map[string][]string           `json:"multiValueQueryStringParameters"`
	PathParameters                  map[string]string             `json:"pathParameters"`
	StageVariables                  map[string]string             `json:"stageVariables"`
	RequestContext                  APIGatewayProxyRequestContext `json:"requestContext"`
	Body                            string                        `json:"body"`
	IsBase64Encoded                 bool                          `json:"isBase64Encoded,omitempty"`
}
```

API Gateway request を受けて route ごとに usecase へ振り分ける [internal/presentation/user/handler.go](/home/dev/backend/internal/presentation/user/handler.go)
```go
func (h Handler) Handle(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	route, err := parseRoute(req)
	if err != nil {
		return cp.BadRequestResponse(), nil
	}

	switch route.kind {
```
