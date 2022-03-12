package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID       uint64    `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Phrase   string    `json:"phrase,omitempty"`
	CreateAt time.Time `json:"createdAt,omitempty"`
	UpdateAt time.Time `json:"updatedAt,omitempty"`
}

func (user *User) Prepare() error {
	user.format()

	if err := user.validate(); err != nil {
		return err
	}

	return nil
}

func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("o campo nome é obrigatório")
	}

	if user.Email == "" {
		return errors.New("o campo e-mail é obrigatório")
	}

	if user.Nick == "" {
		return errors.New("o campo nick é obrigatório")
	}

	if user.Phrase == "" {
		return errors.New("o campo phrase é obrigatório")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Nick = strings.TrimSpace(user.Nick)
}
