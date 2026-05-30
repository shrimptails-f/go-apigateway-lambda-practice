package application

import (
	"context"
	"errors"
	"testing"

	"backend/internal/user/domain"
)

func TestListUseCaseExecute(t *testing.T) {
	t.Parallel()

	t.Run("returns users from repository", func(t *testing.T) {
		t.Parallel()

		expected := []domain.User{{ID: "user-1", Name: "Taro", Email: "taro@example.com"}}
		repo := stubUserRepository{users: expected}

		useCase := NewListUseCase(repo)

		actual, err := useCase.Execute(context.Background())
		if err != nil {
			t.Fatalf("Execute() error = %v", err)
		}
		if len(actual) != len(expected) || actual[0] != expected[0] {
			t.Fatalf("Execute() = %#v, want %#v", actual, expected)
		}
	})

	t.Run("returns repository error", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("list failed")
		repo := stubUserRepository{listErr: expectedErr}

		useCase := NewListUseCase(repo)

		_, err := useCase.Execute(context.Background())
		if !errors.Is(err, expectedErr) {
			t.Fatalf("Execute() error = %v, want %v", err, expectedErr)
		}
	})
}

type stubUserRepository struct {
	users      []domain.User
	listErr    error
	detailUser domain.User
	detailErr  error
}

func (s stubUserRepository) List(context.Context) ([]domain.User, error) {
	return s.users, s.listErr
}

func (s stubUserRepository) GetByID(context.Context, string) (domain.User, error) {
	return s.detailUser, s.detailErr
}
