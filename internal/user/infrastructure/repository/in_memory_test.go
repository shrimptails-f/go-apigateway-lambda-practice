package repository

import (
	"context"
	"errors"
	"testing"

	"backend/internal/user/domain"
)

func TestInMemoryUserRepositoryList(t *testing.T) {
	t.Parallel()

	repo := NewInMemoryUserRepository()

	users, err := repo.List(context.Background())
	if err != nil {
		t.Fatalf("List() error = %v", err)
	}
	if len(users) != 3 {
		t.Fatalf("List() len = %d, want 3", len(users))
	}
}

func TestInMemoryUserRepositoryGetByID(t *testing.T) {
	t.Parallel()

	repo := NewInMemoryUserRepository()

	t.Run("returns matching user", func(t *testing.T) {
		t.Parallel()

		user, err := repo.GetByID(context.Background(), "user-1")
		if err != nil {
			t.Fatalf("GetByID() error = %v", err)
		}
		if user.ID != "user-1" {
			t.Fatalf("GetByID() ID = %s, want user-1", user.ID)
		}
	})

	t.Run("returns not found", func(t *testing.T) {
		t.Parallel()

		_, err := repo.GetByID(context.Background(), "missing")
		if !errors.Is(err, domain.ErrUserNotFound) {
			t.Fatalf("GetByID() error = %v, want %v", err, domain.ErrUserNotFound)
		}
	})
}
