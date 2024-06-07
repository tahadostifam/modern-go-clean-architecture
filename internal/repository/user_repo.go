package repository

import "github.com/tahadostifam/modern-go-clean-architecture/internal/model"

type UserRepository interface {
	Create(name string) (*model.User, error)
}
