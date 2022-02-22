package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
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

	error = user.format(step)
	if error != nil {
		return error
	}

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

	error := checkmail.ValidateFormat(user.Email)
	if error != nil {
		return errors.New("Email is invalid")
	}

	if step == "register" && user.Password == "" {
		return errors.New("Password can not be empty")
	}

	return nil
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if step == "register" {
		hash, error := security.Hash(user.Password)
		if error != nil {
			return error
		}

		user.Password = string(hash)
	}

	return nil
}
