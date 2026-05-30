package application

import (
	"context"

	"backend/internal/user/domain"
)

type UserRepository interface {
	List(context.Context) ([]domain.User, error)
	GetByID(context.Context, string) (domain.User, error)
}
