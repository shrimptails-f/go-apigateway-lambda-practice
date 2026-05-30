package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	up "backend/internal/presentation/user"
	"backend/internal/user/application"
	"backend/internal/user/infrastructure/repository"
)

func main() {
	repo := repository.NewInMemoryUserRepository()
	handler := up.NewHandler(
		application.NewListUseCase(repo),
		application.NewGetDetailUseCase(repo),
	)

	lambda.Start(handler.Handle)
}
