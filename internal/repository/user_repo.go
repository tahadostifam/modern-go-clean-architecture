package repository

import (
	"context"

	"github.com/tahadostifam/modern-go-clean-architecture/internal/model"
)

type UserRepository interface {
	GetById(ctx context.Context, id string) (model.User, error)
	Create(ctx context.Context, user model.User) (model.User, error)
	UpdateById(ctx context.Context, id string, user model.User) (model.User, error)
	DeleteById(ctx context.Context, id string) error
	// Additional methods related to users...
}
type UserRedisRepository interface {
	GetById(id string) (map[string]string, error)
	Create(user model.User) error
	DeleteById(id string) error
	// Additional methods related to users...
}
