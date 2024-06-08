package user_service

import (
	"context"

	"github.com/tahadostifam/modern-go-clean-architecture/internal/model"
	"github.com/tahadostifam/modern-go-clean-architecture/internal/repository"
)

type Service interface {
	Create(ctx context.Context, name, email, password string) (*model.User, error)
}

type service struct {
	// repos
	// other services

	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) (Service, error) {
	return &service{userRepo}, nil
}

func (s *service) Create(ctx context.Context, name string, email string, password string) (*model.User, error) {
	passwordHash := password + "hash me =/"

	return s.userRepo.Create(name, email, passwordHash)
}
