package services

import (
	"github.com/Ira11111/todo-app/internal/models"
	"github.com/Ira11111/todo-app/internal/repository"
)

type Authorization interface {
	generatePasswordHash(password string) string
	Register(user models.User) (int, error)

	GenerateToken(name string) (string, error)
}

type List interface {
}

type Task interface {
}

type Service struct {
	Authorization
	List
	Task
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
