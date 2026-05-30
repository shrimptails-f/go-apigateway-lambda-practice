package presentation

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"

	commonpresentation "backend/internal/common/presentation"
)

func okResponse(payload any) (events.APIGatewayProxyResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return commonpresentation.InternalServerErrorResponse(), nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
		Headers:    jsonHeaders(),
	}, nil
}

func jsonHeaders() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
	}
}
