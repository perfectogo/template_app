package service

import (
	template "github.com/perfectogo/template_app"
	"github.com/perfectogo/template_app/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user template.User) (int, error)
	/* GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error) */
}

type TodoList interface {
	/* Create(userId int, list template.TodoList) (int, error)
	GetAll(userId int) ([]template.TodoList, error)
	GetById(userId, listId int) (template.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input template.UpdateListInput) error */
}

type TodoItem interface {
	/* Create(userId, listId int, item template.TodoItem) (int, error)
	GetAll(userId, listId int) ([]template.TodoItem, error)
	GetById(userId, itemId int) (template.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input template.UpdateItemInput) error */
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		/* TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList), */
	}
}
