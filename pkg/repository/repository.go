package repository

import (
	"github.com/jmoiron/sqlx"
	template "github.com/perfectogo/template_app"
)

type Authorization interface {
	CreateUser(user template.User) (int, error)
	GetUser(username, password string) (template.User, error)
}

type TodoList interface {
	/* Create(userId int, list template.TodoList) (int, error)
	GetAll(userId int) ([]template.TodoList, error)
	GetById(userId, listId int) (template.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input template.UpdateListInput) error
	*/
}

type TodoItem interface {
	/* Create(listId int, item template.TodoItem) (int, error)
	GetAll(userId, listId int) ([]template.TodoItem, error)
	GetById(userId, itemId int) (template.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input template.UpdateItemInput) error */
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		/* TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db), */
	}
}
