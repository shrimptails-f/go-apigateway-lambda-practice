package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"backend/internal/user/application"
	"backend/internal/user/infrastructure/repository"
	"backend/internal/user/presentation"
)

func main() {
	repo := repository.NewInMemoryUserRepository()
	handler := presentation.NewHandler(
		application.NewListUseCase(repo),
		application.NewGetDetailUseCase(repo),
	)

	lambda.Start(handler.Handle)
}
