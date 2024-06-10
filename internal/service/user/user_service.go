package user_service

import (
	"context"
	"time"

	"github.com/tahadostifam/modern-go-clean-architecture/internal/model"
	"github.com/tahadostifam/modern-go-clean-architecture/internal/repository"
)

type Service interface {
	Get(ctx context.Context,id string)(model.User,error)
	Create(ctx context.Context,fname,lname,username,email,phoneNumber,password string) (model.User, error)
	Update(ctx context.Context,id,fname,lname,username,email,phoneNumber,password string) (model.User, error)
	Delete(ctx context.Context,id string)error
}

type service struct {
	// repos
	// other services

	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) (Service, error) {
	return &service{userRepo}, nil
}
func (s *service) Get(ctx context.Context,id string) (model.User, error) {
	return s.userRepo.GetById(ctx,id)
}
func (s *service) Create(ctx context.Context,fname,lname,username,email,phoneNumber,password string) (model.User, error) {
	var user model.User
	hashedPassword := password + "hash me =/"
	user.First_name = fname
	user.Last_name = lname
	user.Username = username
	user.Email = email
	user.Phone_number = phoneNumber
	user.Password = hashedPassword
	user.CreatedAt = time.Now().Local().UTC()
	user.UpdatedAt = time.Now().Local().UTC()
	return s.userRepo.Create(ctx,user)
}
func (s *service) Update(ctx context.Context,id,fname,lname,username,email,phoneNumber,password string) (model.User, error) {
	user,err := s.Get(ctx,id)
	if err != nil{
		return user,err
	}
	hashedPassword := password + "hash me =/"
	user.First_name = fname
	user.Last_name = lname
	user.Username = username
	user.Email = email
	user.Phone_number = phoneNumber
	user.Password = hashedPassword
	user.UpdatedAt = time.Now().Local().UTC()
	return s.userRepo.UpdateById(ctx,id,user)
}
func (s *service) Delete(ctx context.Context,id string) error {
	return s.userRepo.DeleteById(ctx,id)
}
