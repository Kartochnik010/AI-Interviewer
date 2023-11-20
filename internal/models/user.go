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
	Username                    string
	YearsOfCommercialExperience string
	CurrentPosition             string
	DesiredPosition             string
	Stack                       string
	Messages                    []Message
}
type UserResponse struct {
	Username string
	Content  string
}

type UserModel struct {
	DB *sql.DB
}

func (u *UserModel) Insert(user User) error {
	log.Printf("Insert " + user.Username)
	return nil
}

func (u *UserModel) Get(id string) (*User, error) {
	log.Printf("Get user with id " + id)
	return nil, nil
}
