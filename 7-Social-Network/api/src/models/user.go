package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID           uint64    `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	Nick         string    `json:"nick,omitempty"`
	Email        string    `json:"email,omitempty"`
	Password     string    `json:"password,omitempty"`
	CreationDate time.Time `json:"creationDate,omitempty"`
}

func (user *User) Prepare(step string) error {
	error := user.validate(step)
	if error != nil {
		return error
	}

	user.format()

	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("Name can not be empty")
	}

	if user.Nick == "" {
		return errors.New("Nick can not be empty")
	}

	if user.Email == "" {
		return errors.New("Email can not be empty")
	}

	if step == "register" && user.Password == "" {
		return errors.New("Password can not be empty")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
