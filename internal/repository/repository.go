package repository

import (
	"github.com/Ira11111/todo-app/internal/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
}

type List interface {
}

type Task interface {
}

type Repository struct {
	Authorization
	List
	Task
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
	}
}
