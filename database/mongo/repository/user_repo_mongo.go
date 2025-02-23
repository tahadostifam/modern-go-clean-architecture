package repository_mongo

import (
	"context"

	db_mongo "github.com/tahadostifam/modern-go-clean-architecture/database/mongo"
	"github.com/tahadostifam/modern-go-clean-architecture/internal/model"
	"github.com/tahadostifam/modern-go-clean-architecture/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	Users *mongo.Collection
}

func NewUserMongoRepository(mongoClient *mongo.Database) repository.UserRepository {
	return &repo{Users: mongoClient.Collection(db_mongo.UsersCollection)}
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
