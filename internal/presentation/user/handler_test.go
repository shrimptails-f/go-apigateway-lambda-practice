package presentation

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/aws/aws-lambda-go/events"

	"backend/internal/user/domain"
)

func TestHandlerHandle(t *testing.T) {
	t.Parallel()

	t.Run("returns users for list route", func(t *testing.T) {
		t.Parallel()

		handler := NewHandler(
			stubListUseCase{users: []domain.User{{ID: "user-1", Name: "Taro", Email: "taro@example.com"}}},
			stubGetDetailUseCase{},
		)

		resp, err := handler.Handle(context.Background(), events.APIGatewayProxyRequest{
			HTTPMethod: "GET",
			Path:       "/users",
		})
		if err != nil {
			t.Fatalf("Handle() error = %v", err)
		}
		if resp.StatusCode != 200 {
			t.Fatalf("Handle() status = %d, want 200", resp.StatusCode)
		}

		var users []domain.User
		if unmarshalErr := json.Unmarshal([]byte(resp.Body), &users); unmarshalErr != nil {
			t.Fatalf("json.Unmarshal() error = %v", unmarshalErr)
		}
		if len(users) != 1 || users[0].ID != "user-1" {
			t.Fatalf("Handle() body = %#v, want user-1", users)
		}
	})

	t.Run("returns user detail", func(t *testing.T) {
		t.Parallel()

		handler := NewHandler(
			stubListUseCase{},
			stubGetDetailUseCase{user: domain.User{ID: "user-1", Name: "Taro", Email: "taro@example.com"}},
		)

		resp, err := handler.Handle(context.Background(), events.APIGatewayProxyRequest{
			HTTPMethod: "GET",
			Path:       "/user/user-1",
		})
		if err != nil {
			t.Fatalf("Handle() error = %v", err)
		}
		if resp.StatusCode != 200 {
			t.Fatalf("Handle() status = %d, want 200", resp.StatusCode)
		}
	})

	t.Run("returns bad request for invalid route", func(t *testing.T) {
		t.Parallel()

		handler := NewHandler(stubListUseCase{}, stubGetDetailUseCase{})

		resp, err := handler.Handle(context.Background(), events.APIGatewayProxyRequest{
			HTTPMethod: "GET",
			Path:       "/user/a/b",
		})
		if err != nil {
			t.Fatalf("Handle() error = %v", err)
		}
		if resp.StatusCode != 400 {
			t.Fatalf("Handle() status = %d, want 400", resp.StatusCode)
		}
	})

	t.Run("returns not found for missing user", func(t *testing.T) {
		t.Parallel()

		handler := NewHandler(
			stubListUseCase{},
			stubGetDetailUseCase{err: domain.ErrUserNotFound},
		)

		resp, err := handler.Handle(context.Background(), events.APIGatewayProxyRequest{
			HTTPMethod: "GET",
			Path:       "/user/missing",
		})
		if err != nil {
			t.Fatalf("Handle() error = %v", err)
		}
		if resp.StatusCode != 404 {
			t.Fatalf("Handle() status = %d, want 404", resp.StatusCode)
		}
	})

	t.Run("returns internal server error when usecase fails", func(t *testing.T) {
		t.Parallel()

		handler := NewHandler(
			stubListUseCase{err: errors.New("boom")},
			stubGetDetailUseCase{},
		)

		resp, err := handler.Handle(context.Background(), events.APIGatewayProxyRequest{
			HTTPMethod: "GET",
			Path:       "/users",
		})
		if err != nil {
			t.Fatalf("Handle() error = %v", err)
		}
		if resp.StatusCode != 500 {
			t.Fatalf("Handle() status = %d, want 500", resp.StatusCode)
		}
	})
}

type stubListUseCase struct {
	users []domain.User
	err   error
}

func (s stubListUseCase) Execute(context.Context) ([]domain.User, error) {
	return s.users, s.err
}

type stubGetDetailUseCase struct {
	user domain.User
	err  error
}

func (s stubGetDetailUseCase) Execute(context.Context, string) (domain.User, error) {
	return s.user, s.err
}
