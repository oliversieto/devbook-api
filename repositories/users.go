package repositories

import (
	"database/sql"
	"devbook-api/src/models"
)

type users struct {
	database *sql.DB
}

func NewUserRepository(database *sql.DB) *users {
	return &users{database}
}

func (repository users) Create(user models.User) (uint64, error) {
	statement, err := repository.database.Prepare("insert into users (name, nick, email, phrase) values (?, ?, ?, ?)")

	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := repository.database.Exec(user.Name, user.Nick, user.Email, user.Phrase)

	if err != nil {
		return 0, err
	}

	ID, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(ID), nil
}
