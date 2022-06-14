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

func (repository Publications) GetByID(ID uint64) (models.Publication, error) {
	row, err := repository.db.Query(`
		select p.*, u.nick
		from publications p inner join users u on u.id = p.author_id
		where p.id = ?
	`, ID)
	if err != nil {
		return models.Publication{}, err
	}
	defer row.Close()

	var publication models.Publication

	if row.Next() {
		err = row.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreationDate,
			&publication.AuthorNick)
		if err != nil {
			return models.Publication{}, err
		}
	}

	return publication, nil
}

func (repository Publications) Get(userID uint64) ([]models.Publication, error) {
	rows, err := repository.db.Query(`
		select p.*, u.nick
		from publications p
		inner join users u on u.id = p.author_id
		where p.author_id = ?
		UNION
		select p.*, u.nick
		from publications p
		inner join users u on u.id = p.author_id
		inner join followers f on p.author_id = f.user_id
		where f.follower_id = ?
		order by 1 desc
	`, userID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var publications []models.Publication

	for rows.Next() {
		var publication models.Publication

		err = rows.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreationDate,
			&publication.AuthorNick)
		if err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

func (repository Publications) Update(publicationID uint64, publication models.Publication) error {
	statement, err := repository.db.Prepare(
		"update publications set title = ?, content = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(publication.Title, publication.Content, publicationID)
	if err != nil {
		return err
	}

	return nil
}

func (repository Publications) Delete(publicationID uint64) error {
	statement, err := repository.db.Prepare(
		"delete from publications where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(publicationID)
	if err != nil {
		return err
	}

	return nil
}

func (repository Publications) GetPublicationsByUser(userID uint64) ([]models.Publication, error) {
	rows, err := repository.db.Query(`
		select p.*, u.nick
		from publications p inner join users u on u.id = p.author_id
		where p.author_id = ?
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var publications []models.Publication

	for rows.Next() {
		var publication models.Publication

		err = rows.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreationDate,
			&publication.AuthorNick)
		if err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

func (repository Publications) Like(publicationID uint64) error {
	statement, err := repository.db.Prepare(
		"update publications set likes = likes + 1 where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(publicationID)
	if err != nil {
		return err
	}

	return nil
}

func (repository Publications) Dislike(publicationID uint64) error {
	statement, err := repository.db.Prepare(`
		update publications set
		likes = case when likes > 0 then likes - 1 else likes end
		where id = ?
		`,
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(publicationID)
	if err != nil {
		return err
	}

	return nil
}
