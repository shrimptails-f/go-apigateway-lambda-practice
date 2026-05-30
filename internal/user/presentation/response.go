package presentation

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func okResponse(payload any) (events.APIGatewayProxyResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return internalServerErrorResponse(), nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
		Headers:    jsonHeaders(),
	}, nil
}

func badRequestResponse() events.APIGatewayProxyResponse {
	return errorResponse(http.StatusBadRequest, "invalid request path")
}

func notFoundResponse() events.APIGatewayProxyResponse {
	return errorResponse(http.StatusNotFound, "resource not found")
}

func internalServerErrorResponse() events.APIGatewayProxyResponse {
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
