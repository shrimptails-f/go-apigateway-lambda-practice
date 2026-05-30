package presentation

import (
	"errors"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestParseRoute(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		request   events.APIGatewayProxyRequest
		wantRoute route
		wantErr   error
	}{
		{
			name:      "list users",
			request:   events.APIGatewayProxyRequest{HTTPMethod: "GET", Path: "/users"},
			wantRoute: route{kind: routeListUsers},
		},
		{
			name:      "get user detail",
			request:   events.APIGatewayProxyRequest{HTTPMethod: "GET", Path: "/user/user-1"},
			wantRoute: route{kind: routeGetUserDetail, userID: "user-1"},
		},
		{
			name:    "invalid nested user route",
			request: events.APIGatewayProxyRequest{HTTPMethod: "GET", Path: "/user/a/b"},
			wantErr: errInvalidRoute,
		},
		{
			name:    "unsupported method returns empty route",
			request: events.APIGatewayProxyRequest{HTTPMethod: "POST", Path: "/users"},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual, err := parseRoute(tt.request)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("parseRoute() error = %v, want %v", err, tt.wantErr)
			}
			if actual != tt.wantRoute {
				t.Fatalf("parseRoute() = %#v, want %#v", actual, tt.wantRoute)
			}
		})
	}
}
