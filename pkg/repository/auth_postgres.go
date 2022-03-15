package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	template "github.com/perfectogo/template_app"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
func (r *AuthPostgres) CreateUser(user template.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING user_id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
func (r *AuthPostgres) GetUser(username, password string) (template.User, error) {
	var user template.User
	//var id int
	query := fmt.Sprintf("SELECT user_id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user.Id, query, username, password)
	log.Println(user)
	return user, err
}
