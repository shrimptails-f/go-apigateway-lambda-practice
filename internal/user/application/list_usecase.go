package application

import (
	"context"

	"backend/internal/user/domain"
)

type ListUseCase struct {
	repository UserRepository
}

func NewListUseCase(repository UserRepository) *ListUseCase {
	return &ListUseCase{repository: repository}
}

func (u ListUseCase) Execute(ctx context.Context) ([]domain.User, error) {
	return u.repository.List(ctx)
}
