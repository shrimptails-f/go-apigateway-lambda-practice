package presentation

import (
	"context"

	"github.com/aws/aws-lambda-go/events"

	"backend/internal/user/domain"
)

type Handler struct {
	listUseCase      ListUseCase
	getDetailUseCase GetDetailUseCase
}

func NewHandler(listUseCase ListUseCase, getDetailUseCase GetDetailUseCase) Handler {
	return Handler{
		listUseCase:      listUseCase,
		getDetailUseCase: getDetailUseCase,
	}
}

func (h Handler) Handle(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	route, err := parseRoute(req)
	if err != nil {
		return badRequestResponse(), nil
	}

	switch route.kind {
	case routeListUsers:
		users, execErr := h.listUseCase.Execute(ctx)
		if execErr != nil {
			return internalServerErrorResponse(), nil
		}

		return okResponse(users)
	case routeGetUserDetail:
		user, execErr := h.getDetailUseCase.Execute(ctx, route.userID)
		if execErr != nil {
			if execErr == domain.ErrUserNotFound {
				return notFoundResponse(), nil
			}

			return internalServerErrorResponse(), nil
		}

		return okResponse(user)
	default:
		return notFoundResponse(), nil
	}
}
