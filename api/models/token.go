package models

import (
	"database/sql"
	"fmt"
	"log"
	"urlshorter/utils"
)

var TOKEN_TABLE = "tokens"

type Token struct {
	ID        int    `json:"id"`
	UserId    string `json:"user_id"`
	Token     string `json:"token"`
	CreatedAt string `json:"created_at"`
}

type TokenModel struct {
	DB *sql.DB
}

func (m *TokenModel) Create(userId int) (string, int) {
	token := utils.RandomString(64)
	query := fmt.Sprintf("INSERT INTO %s (user_id, token) VALUES(?, ?)", TOKEN_TABLE)

	res, err := m.DB.Exec(query, userId, token)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return token, int(id)
}

func (m *TokenModel) GetUser(token string) User {
	query := fmt.Sprintf("SELECT %s.id, %s.name FROM %s INNER JOIN %s ON %s.id = %s.user_id WHERE %s.token = ? LIMIT 1", USER_TABLE, USER_TABLE, TOKEN_TABLE, TOKEN_TABLE, USER_TABLE, USER_TABLE, TOKEN_TABLE, TOKEN_TABLE)

	row := m.DB.QueryRow(query, token)

	user := User{}
	row.Scan(
		&user.ID,
		&user.Name,
	)

	return user
}
