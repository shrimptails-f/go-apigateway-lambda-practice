package application

import (
	"context"

	"backend/internal/user/domain"
)

type GetDetailUseCase struct {
	repository UserRepository
}

func NewGetDetailUseCase(repository UserRepository) *GetDetailUseCase {
	return &GetDetailUseCase{repository: repository}
}

func (u GetDetailUseCase) Execute(ctx context.Context, id string) (domain.User, error) {
	return u.repository.GetByID(ctx, id)
}
