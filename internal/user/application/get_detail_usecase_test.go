package application

import (
	"context"
	"errors"
	"testing"

	"backend/internal/user/domain"
)

func TestGetDetailUseCaseExecute(t *testing.T) {
	t.Parallel()

	t.Run("returns user from repository", func(t *testing.T) {
		t.Parallel()

		expected := domain.User{ID: "user-1", Name: "Taro", Email: "taro@example.com"}
		repo := stubUserRepository{detailUser: expected}

		useCase := NewGetDetailUseCase(repo)

		actual, err := useCase.Execute(context.Background(), "user-1")
		if err != nil {
			t.Fatalf("Execute() error = %v", err)
		}
		if actual != expected {
			t.Fatalf("Execute() = %#v, want %#v", actual, expected)
		}
	})

	t.Run("returns repository error", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("detail failed")
		repo := stubUserRepository{detailErr: expectedErr}

		useCase := NewGetDetailUseCase(repo)

		_, err := useCase.Execute(context.Background(), "user-1")
		if !errors.Is(err, expectedErr) {
			t.Fatalf("Execute() error = %v, want %v", err, expectedErr)
		}
	})
}
