package models

import (
	"database/sql"
	"log"
)

// Статистика
// - показать юзеру когда был его первый запрос,
// - сколько всего запросов было
// - можно добавить ещё показателей, которые мы можем получить из хранимых данных.

type User struct {
	ID              int
	Name            string
	TechStack       []string
	CurrentPosition Position
	GoalPosition    Position
	About           string
	CreatedAt       string
}
type Position string

type UserModel struct {
	DB *sql.DB
}

func (u *UserModel) Insert(user User) error {
	log.Printf("Insert " + user.Name)
	return nil
}

func (u *UserModel) Get(id string) (*User, error) {
	log.Printf("Get user with id " + id)
	return nil, nil
}
