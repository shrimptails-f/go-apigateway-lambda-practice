package repository

import (
	"context"

	"backend/internal/user/domain"
)

type InMemoryUserRepository struct {
	users []domain.User
}

func NewInMemoryUserRepository() InMemoryUserRepository {
	return InMemoryUserRepository{
		users: []domain.User{
			{ID: "user-1", Name: "Taro Yamada", Email: "taro@example.com"},
			{ID: "user-2", Name: "Hanako Sato", Email: "hanako@example.com"},
			{ID: "user-3", Name: "John Suzuki", Email: "john@example.com"},
		},
	}
}

func (r InMemoryUserRepository) List(context.Context) ([]domain.User, error) {
	return r.users, nil
}

func (r InMemoryUserRepository) GetByID(_ context.Context, id string) (domain.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}

	return domain.User{}, domain.ErrUserNotFound
}
