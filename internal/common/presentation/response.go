package presentation

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func BadRequestResponse() events.APIGatewayProxyResponse {
	return errorResponse(http.StatusBadRequest, "invalid request path")
}

func NotFoundResponse() events.APIGatewayProxyResponse {
	return errorResponse(http.StatusNotFound, "resource not found")
}

func InternalServerErrorResponse() events.APIGatewayProxyResponse {
	return errorResponse(http.StatusInternalServerError, "internal server error")
}

func errorResponse(statusCode int, message string) events.APIGatewayProxyResponse {
	body, err := json.Marshal(map[string]string{"message": message})
	if err != nil {
		body = []byte(`{"message":"internal server error"}`)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       string(body),
		Headers:    jsonHeaders(),
	}
}

func jsonHeaders() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
	}
}
