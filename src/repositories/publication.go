package repositories

import (
	"api/src/models"
	"database/sql"
)

type publication struct {
	db *sql.DB
}

func NewPublicationRepository(db *sql.DB) *publication {
	return &publication{db}
}

func (repository publication) Create(publication models.Publication) (uint64, error) {
	stmt, err := repository.db.Prepare(
		"insert into publications (author_id, title, content) values (?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(publication.AuthorID, publication.Title, publication.Content)
	if err != nil {
		return 0, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedID), nil
}
