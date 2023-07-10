package migrations

import (
	"urlshorter/utils"
)

var (
	reload bool = false
)

func ApplyMigrations() {
	linksTable()
	userTable()
	authTokensTable()
}

func userTable() {
	db := utils.GetDBConnection()

	defer db.Close()

	if reload {
		query := "DROP TABLE users"

		db.Exec(query)
	}

	query := `
		CREATE TABLE IF NOT EXISTS users (
			id INT PRIMARY KEY AUTO_INCREMENT,
			name VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL
		);
	`

	db.Exec(query)
}

func linksTable() {
	db := utils.GetDBConnection()

	defer db.Close()

	if reload {
		query := "DROP TABLE links"

		db.Exec(query)
	}

	query := `
		CREATE TABLE IF NOT EXISTS links (
			id INT PRIMARY KEY AUTO_INCREMENT,
			link TEXT NOT NULL,
			slug VARCHAR(255) NOT NULL,
		    user_id INT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`

	db.Exec(query)
}

func authTokensTable() {
	db := utils.GetDBConnection()

	defer db.Close()

	if reload {
		query := "DROP TABLE tokens"

		db.Exec(query)
	}

	query := `
		CREATE TABLE IF NOT EXISTS tokens (
			id INT PRIMARY KEY AUTO_INCREMENT,
			user_id INT NOT NULL,
			token VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`

	db.Exec(query)
}
