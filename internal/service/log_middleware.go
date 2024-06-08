package service

import (
	"context"
	"time"
)

// LogMiddleware also implement UserService interface
type LogMiddleware struct {
	service Service
}

func NewLogMiddleware(service Service) *LogMiddleware {
	return &LogMiddleware{
		service: service,
	}
}

func (m *LogMiddleware) Create(ctx context.Context, req createUserRequest) (createUserResponse, error) {
	// log request and response here
	defer func(begin time.Time) {
		// log time spend, request, response here, ex: time.Since(begin) -> time spend for the service to response to the request
	}(time.Now())

	return m.service.Create(ctx, req)
}
