package models

import (
	"errors"
	"strings"
	"time"
)

type Publication struct {
	ID         uint64    `json:"id,omitempty"`
	AuthorID   uint64    `json:"author_id,omitempty"`
	AuthorNick string    `json:"author_nick,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}

func (p *Publication) Prepare() error {
	if err := p.validate(); err != nil {
		return err
	}

	err := p.format()

	if err != nil {
		return err
	}

	return nil
}

func (p *Publication) validate() error {
	if p.AuthorID == 0 {
		return errors.New("author_id is required")
	}

	if p.Title == "" {
		return errors.New("title is required")
	}

	if p.Content == "" {
		return errors.New("content is required")
	}

	return nil
}

func (p *Publication) format() error {
	p.Title = strings.TrimSpace(p.Title)
	p.Content = strings.TrimSpace(p.Content)

	return nil
}
