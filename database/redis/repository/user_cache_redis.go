package repository_redis

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
	redisDB "github.com/tahadostifam/modern-go-clean-architecture/database/redis"
	"github.com/tahadostifam/modern-go-clean-architecture/internal/model"
	"github.com/tahadostifam/modern-go-clean-architecture/internal/repository"
)

type UserRedisRepository struct {
	RCli *redis.Client
}
func NewUserRedisRepository()*UserRedisRepository{
	return &UserRedisRepository{
		RCli:redisDB.NewRedisConnection() ,
	}
}
//just for ensure
var _UserRedisRepository repository.UserRedisRepository = NewUserRedisRepository()
func (ur *UserRedisRepository) GetById(id string) (map[string]string, error) {
	exists := ur.RCli.Exists(context.Background(), fmt.Sprintf("user:%s", id))
	if exists.Val() == 0 {
		return nil, errors.New("users not found")
	}
	redisMapRes := ur.RCli.HGetAll(context.Background(), fmt.Sprintf("user:%s", id))
	if redisMapRes.Err() != nil {
		return nil, errors.New("users not found")
	}
	return redisMapRes.Val(),nil
}
func (ur *UserRedisRepository) Create(user model.User) error {
	redisRes := ur.RCli.HMSet(context.Background(), fmt.Sprintf("user:%d", user.ID), map[string]interface{}{
		"first_name":   user.First_name,
		"last_name":    user.Last_name,
		"username":    user.Username,
		"email":       user.Email,
		"phone_number": user.Phone_number,
		"createdAt":   user.CreatedAt,
		"updatedAt":   user.UpdatedAt,
	})
	return redisRes.Err()
}
func (ur *UserRedisRepository) DeleteById(id string) error {
	redisRes := ur.RCli.Del(context.Background(), fmt.Sprintf("user:%s", id))
	return redisRes.Err()
}
