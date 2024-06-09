package repository_mongo

import (
	"context"

	mongoDB "github.com/tahadostifam/modern-go-clean-architecture/database/mongo"
	"github.com/tahadostifam/modern-go-clean-architecture/internal/model"
	"github.com/tahadostifam/modern-go-clean-architecture/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongoRepository struct {
	Users *mongo.Collection
}

func NewUserMongoRepository() *UserMongoRepository {
	return &UserMongoRepository{
		//TODO:get database name and collection names from configs
		Users: mongoDB.NewMongoConnection().Database("mydb").Collection("users"),
	}
}

// just for ensure
var _UserRepo repository.UserRepository = NewUserMongoRepository()

func (ur *UserMongoRepository) GetById(ctx context.Context, id string) (model.User, error) {
	var u model.User
	filter := bson.D{primitive.E{Key: "_id", Value: id}}

	err := ur.Users.FindOne(ctx, filter).Decode(&u)
	if err != nil {
		return model.User{}, err
	}
	return u, nil
}
func (ur *UserMongoRepository) Create(ctx context.Context, user model.User) (model.User, error) {
	_, err := ur.Users.InsertOne(ctx, &user)
	return user, err
}
func (ur *UserMongoRepository) UpdateById(ctx context.Context, id string, user model.User) (model.User, error) {
	_, err := ur.Users.UpdateByID(ctx, id, &user)
	return user, err
}
func (ur *UserMongoRepository) DeleteById(ctx context.Context, id string) error {
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	_, err := ur.Users.DeleteOne(ctx, filter)
	return err
}
