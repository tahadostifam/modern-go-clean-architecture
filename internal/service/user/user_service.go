package user_service

import (
	"context"

	"github.com/tahadostifam/modern-go-clean-architecture/internal/model"
	"github.com/tahadostifam/modern-go-clean-architecture/internal/repository"
)

type UserService interface {
	Get(ctx context.Context, id string) (model.User, error)
	Create(ctx context.Context, fname, lname, username, email, phoneNumber, password string) (model.User, error)
	Update(ctx context.Context, id, fname, lname, username, email, phoneNumber, password string) (model.User, error)
	Delete(ctx context.Context, id string) error
}

type service struct {
	userRepo repository.UserRepository
}

func NewService(userRepo repository.UserRepository) (UserService, error) {
	return &service{userRepo}, nil
}

func (s *service) Create(ctx context.Context, fname string, lname string, username string, email string, phoneNumber string, password string) (model.User, error) {
	panic("unimplemented")
}

func (s *service) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

func (s *service) Get(ctx context.Context, id string) (model.User, error) {
	panic("unimplemented")
}

func (s *service) Update(ctx context.Context, id string, fname string, lname string, username string, email string, phoneNumber string, password string) (model.User, error) {
	panic("unimplemented")
}
