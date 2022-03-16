package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	template "github.com/perfectogo/template_app"
	"github.com/sirupsen/logrus"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userId int, list template.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	row := tx.QueryRow(
		`insert into users_lists (title, description) value ($1, $2) returning id`,
		list.Title,
		list.Description,
	)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	_, err = tx.Exec(
		`insert into users_lists (user_id, list_id) values ($1, $2)`,
		userId,
		id,
	)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *TodoListPostgres) GetAll(userId int) ([]template.TodoList, error) {
	var lists []template.TodoList

	err := r.db.Select(
		&lists,
		`SELECT tl.id, tl.title, tl.description FROM todo_lists tl INNER JOIN users_lists ul on tl.id = ul.list_id WHERE ul.user_id = $1`,
		userId)

	return lists, err
}

func (r *TodoListPostgres) GetById(userId, listId int) (template.TodoList, error) {
	var list template.TodoList

	err := r.db.Get(
		&list,
		`select tl.id, tl.title, tl.description from todo_lists tl INNER JOIN users_lists ul on tl.id = ul.list_id where ul.user_id = $1 and ul.list_id = $2`,
		userId,
		listId,
	)

	return list, err
}

func (r *TodoListPostgres) Delete(userId, listId int) error {
	_, err := r.db.Exec(
		`delete from todo_lists tl using users_lists ul where tl.id = ul.list_id and ul.user_id=$1 and ul.list_id=$2`,
		userId,
		listId,
	)

	return err
}

func (r *TodoListPostgres) Update(userId, listId int, input template.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		todoListsTable, setQuery, usersListsTable, argId, argId+1)
	args = append(args, listId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
