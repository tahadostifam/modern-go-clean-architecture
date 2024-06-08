package service

import "context"

type Service interface {
	Create(ctx context.Context, req createUserRequest) (res createUserResponse, err error)
}

// service Implements UserService interface
type service struct {
	// repos
	// other services
}

// NewUserService returns new instance of UserService, and it takes repos instances, some constant settings, other services as the input arguments
func NewUserService() (Service, error) {
	var s Service
	s = &service{}
	// middlewares adds here
	s = NewLogMiddleware(s) // now it's act like a middleware
	// add more middlewares here, ex: validation, sentry, ...
	return s, nil
}

func (s *service) Create(ctx context.Context, req createUserRequest) (createUserResponse, error) {
	return createUserResponse{}, nil
}
