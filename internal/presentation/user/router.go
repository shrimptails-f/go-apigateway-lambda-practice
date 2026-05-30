package presentation

import (
	"errors"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

var errInvalidRoute = errors.New("invalid route")

type routeKind string

const (
	routeListUsers     routeKind = "list_users"
	routeGetUserDetail routeKind = "get_user_detail"
)

type route struct {
	kind   routeKind
	userID string
}

func parseRoute(req events.APIGatewayProxyRequest) (route, error) {
	if req.HTTPMethod != "GET" {
		return route{}, nil
	}

	path := strings.Trim(req.Path, "/")
	if path == "users" {
		return route{kind: routeListUsers}, nil
	}

	parts := strings.Split(path, "/")
	if len(parts) == 2 && parts[0] == "user" && parts[1] != "" {
		return route{kind: routeGetUserDetail, userID: parts[1]}, nil
	}

	if strings.HasPrefix(path, "user/") {
		return route{}, errInvalidRoute
	}

	return route{}, nil
}
