package service

import (
	"context"
	"thoropa/internal/model"
	"thoropa/internal/repository"
)

type LinkService struct {
	repo *repository.LinkRepository
}

func NewLinkService(r *repository.LinkRepository) *LinkService {
	return &LinkService{repo: r}
}

func (s *LinkService) Create(ctx context.Context, link *model.Link) error {
	// 👉 aqui entra regra de negócio depois
	return s.repo.Create(ctx, link)
}
