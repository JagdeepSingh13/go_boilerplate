package service

import (
	"github.com/JagdeepSingh13/go_boilerplate/internal/lib/job"
	"github.com/JagdeepSingh13/go_boilerplate/internal/repository"
	"github.com/JagdeepSingh13/go_boilerplate/internal/server"
)

type Services struct {
	Auth *AuthService
	Job  *job.JobService
}

func NewServices(s *server.Server, repos *repository.Repositories) (*Services, error) {
	authService := NewAuthService(s)

	return &Services{
		Job:  s.Job,
		Auth: authService,
	}, nil
}
