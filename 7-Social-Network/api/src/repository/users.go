package repository

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

func (repository users) Create(user models.User) (uint64, error) {
	statement, error := repository.db.Prepare(
		"insert into users (name, nick, email, password) values(?, ?, ?, ?)",
	)
	if error != nil {
		return 0, error
	}
	defer statement.Close()

	result, error := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if error != nil {
		return 0, error
	}

	lastInsertedID, error := result.LastInsertId()
	if error != nil {
		return 0, error
	}

	return uint64(lastInsertedID), nil
}

func (repository users) Get(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	rows, error := repository.db.Query(
		"select id, name, nick, email, creationDate from users where name like ? or nick like ?",
		nameOrNick,
		nameOrNick,
	)
	if error != nil {
		return nil, error
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		error = rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreationDate)
		if error != nil {
			return nil, error
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository users) GetByID(ID uint64) (models.User, error) {
	row, error := repository.db.Query(
		"select id, name, nick, email, creationDate from users where ID = ?",
		ID,
	)
	if error != nil {
		return models.User{}, error
	}
	defer row.Close()

	var user models.User

	if row.Next() {
		error = row.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreationDate)
		if error != nil {
			return models.User{}, error
		}
	}

	return user, nil
}
