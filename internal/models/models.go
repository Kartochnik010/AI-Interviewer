package models

import "database/sql"

type Models struct {
	Users UserModel
	Gpt   GptModel
}

func NewModels(db *sql.DB) *Models {
	return &Models{
		Users: UserModel{DB: db},
		Gpt:   GptModel{DB: db},
	}
}
