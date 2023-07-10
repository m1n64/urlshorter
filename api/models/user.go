package models

import (
	"database/sql"
	"fmt"
	"log"
)

var USER_TABLE = "users"

type User struct {
	ID       int    `json:"ID"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Get(name string, password string) (User, error) {
	query := fmt.Sprintf("SELECT id, name, password FROM %s WHERE name = ? AND password = MD5(?)", USER_TABLE)

	row := m.DB.QueryRow(query, name, password)

	user := User{}

	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Password,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, nil
		}

		return User{}, err
	}

	return user, nil
}

func (m *UserModel) Create(name string, password string) int {
	query := fmt.Sprintf("INSERT INTO %s (name, password) VALUES(?, MD5(?))", USER_TABLE)

	res, err := m.DB.Exec(query, name, password)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return int(id)
}
