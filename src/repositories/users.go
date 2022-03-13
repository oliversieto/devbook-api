package repositories

import (
	"database/sql"
	"devbook-api/src/models"
	"fmt"
	"time"
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

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Phrase)

	if err != nil {
		return 0, err
	}

	ID, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(ID), nil
}

func (repository users) GetAll(search string) ([]models.User, error) {
	search = fmt.Sprintf("%%%s%%", search)

	result, err := repository.database.Query("select id, name, email, nick, createAt from users where lower(name) like ? or lower(nick) like ?", search, search)
	if err != nil {
		return nil, err
	}

	defer result.Close()

	var users []models.User

	for result.Next() {
		var user models.User

		if err := result.Scan(&user.ID, &user.Name, &user.Email, &user.Nick, &user.CreateAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository users) GetOne(ID uint64) (models.User, error) {
	result, err := repository.database.Query("select id, name, email, nick, createAt from users where id = ?", ID)

	if err != nil {
		return models.User{}, err
	}

	defer result.Close()

	var user models.User

	if result.Next() {
		if err := result.Scan(&user.ID, &user.Name, &user.Email, &user.Nick, &user.CreateAt); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository users) Update(user models.User, ID uint64) error {
	prepare, err := repository.database.Prepare("update users set name = ?, nick = ?, email = ?, updateAt = ? where id = ?")
	updateAt := time.Now()

	if err != nil {
		return err
	}

	defer prepare.Close()

	if _, err = prepare.Exec(user.Name, user.Nick, user.Email, updateAt, ID); err != nil {
		return err
	}

	return nil

}

func (repository users) Delete(ID uint64) error {
	prepare, err := repository.database.Prepare("delete from users where id = ?")

	if err != nil {
		return err
	}

	defer prepare.Close()

	if _, err = prepare.Exec(ID); err != nil {
		return err
	}

	return nil
}
