package main

import (
	"context"
	"thoropa/internal/database"
	"thoropa/internal/handler"
	"thoropa/internal/repository"
	"thoropa/internal/router"
	"thoropa/internal/service"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	ctx := context.Background()

	db := database.NewDynamoClient(ctx)

	// camadas
	repo := repository.NewLinkRepository(db)
	service := service.NewLinkService(repo)
	handler := handler.NewLinkHandler(service)

	r := router.SetupRouter(handler)

	r.Run(":8080")
}
