package models

import (
	"database/sql"
	"fmt"
	"log"
	"urlshorter/utils"
)

var LINKS_TABLE = "links"

type Link struct {
	ID        int    `json:"id"`
	Link      string `json:"link"`
	Slug      string `json:"slug"`
	CreatedAt string `json:"created_at"`
}

type LinkModel struct {
	DB *sql.DB
}

func (m *LinkModel) GetLinkBySlug(slug string) (Link, error) {
	query := fmt.Sprintf("SELECT ID, link, slug, created_at FROM %s WHERE slug = ? LIMIT 1", LINKS_TABLE)

	row := m.DB.QueryRow(query, slug)
	link := Link{}
	err := row.Scan(&link.ID, &link.Link, &link.Slug, &link.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return Link{}, nil
		}

		return Link{}, err
	}

	return link, nil
}

func (m *LinkModel) InsertLink(link string) (string, int) {
	slug := utils.RandomString(8)
	query := fmt.Sprintf("INSERT INTO %s (link, slug) VALUES (?, ?)", LINKS_TABLE)

	res, err := m.DB.Exec(query, link, slug)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return slug, int(id)
}
