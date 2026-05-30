package presentation

import (
	"backend/internal/user/domain"
	"context"
)

type ListUseCase interface {
	Execute(ctx context.Context) ([]domain.User, error)
}

type GetDetailUseCase interface {
	Execute(ctx context.Context, id string) (domain.User, error)
}
