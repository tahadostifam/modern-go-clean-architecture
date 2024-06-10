package repository_redis

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/tahadostifam/modern-go-clean-architecture/internal/model"
	"github.com/tahadostifam/modern-go-clean-architecture/internal/repository"
)

type repo struct {
	redisClient *redis.Client
}

func NewUserRedisRepository(redisClient *redis.Client) repository.UserRepository {
	return &repo{redisClient}
}

func (r *repo) Create(ctx context.Context, user model.User) (model.User, error) {
	panic("unimplemented")
}

func (r *repo) DeleteById(ctx context.Context, id string) error {
	panic("unimplemented")
}

func (r *repo) GetById(ctx context.Context, id string) (model.User, error) {
	panic("unimplemented")
}

func (r *repo) UpdateById(ctx context.Context, id string, user model.User) (model.User, error) {
	panic("unimplemented")
}
