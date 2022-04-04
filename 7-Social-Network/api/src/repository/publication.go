package repository

import (
	"api/src/models"
	"database/sql"
)

type Publications struct {
	db *sql.DB
}

func NewPublicationsRepository(db *sql.DB) *Publications {
	return &Publications{db}
}

func (repository Publications) Create(publication models.Publication) (uint64, error) {
	statement, err := repository.db.Prepare(
		"insert into publications (title, content, author_id) values(?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(publication.Title, publication.Title, publication.AuthorID)
	if err != nil {
		return 0, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedID), nil
}
