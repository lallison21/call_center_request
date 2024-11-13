package service

import "github.com/lallison21/call_center_request/auth_service/internal/repository"

type Service struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}
