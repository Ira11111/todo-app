package services

import "github.com/Ira11111/todo-app/internal/repository"

type Authorization interface {
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
	return &Service{}
}
