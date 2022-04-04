package models

import (
	"errors"
	"strings"
	"time"
)

type Publication struct {
	ID           uint64    `json:"id,omitempty"`
	Title        string    `json:"title,omitempty"`
	Content      string    `json:"content,omitempty"`
	AuthorID     uint64    `json:"authorId,omitempty"`
	AuthorNick   uint64    `json:"authorNick,omitempty"`
	Likes        uint64    `json:"likes"`
	CreationDate time.Time `json:"creationDate,omitempty"`
}

func (publication *Publication) Prepare() error {
	err := publication.validate()
	if err != nil {
		return err
	}

	publication.format()

	return nil
}

func (publication *Publication) format() {
	publication.Title = strings.TrimSpace(publication.Title)

	publication.Content = strings.TrimSpace(publication.Content)
}

func (publication *Publication) validate() error {
	if publication.Title == "" {
		return errors.New("title is mandatory and must not be empty")
	}

	if publication.Content == "" {
		return errors.New("content is mandatory and must not be empty")
	}

	return nil
}
