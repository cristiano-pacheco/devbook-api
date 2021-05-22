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

func (repository publication) Get(ID uint64) (models.Publication, error) {
	p := models.Publication{}
	err := repository.db.QueryRow(
		`select p.id, author_id, u.nick, title, content, likes, p.created_at 
		from publications p
		join users u on p.author_id = u.id
		where p.id = ?`,
		ID,
	).Scan(&p.ID, &p.AuthorID, &p.AuthorNick, &p.Title, &p.Content, &p.Likes, &p.CreatedAt)

	if err != nil {
		return models.Publication{}, err
	}

	return p, nil
}

func (repository publication) GetByUserID(ID uint64) ([]models.Publication, error) {
	rows, err := repository.db.Query(
		`select distinct p.id, author_id, u.nick, title, content, likes, p.created_at 
		from publications p
		join users u on p.author_id = u.id
		join followers f on p.author_id = f.user_id
		where u.id = ? or f.user_id = ? order by 1 desc`,
		ID,
		ID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var publications []models.Publication

	for rows.Next() {
		var p models.Publication
		err := rows.Scan(&p.ID, &p.AuthorID, &p.AuthorNick, &p.Title, &p.Content, &p.Likes, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		publications = append(publications, p)
	}

	return publications, nil
}
