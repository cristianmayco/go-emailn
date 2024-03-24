package campaign

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

type Contatc struct {
	Email string
}

type Campaign struct {
	ID        string
	Name      string
	CreatedOn time.Time
	Content   string
	Contacts  []Contatc
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	if content == "" {
		return nil, errors.New("content is required")
	}

	if len(emails) == 0 {
		return nil, errors.New("contacts is required")
	}

	constatcs := make([]Contatc, len(emails))

	for index, email := range emails {
		constatcs[index].Email = email
	}

	return &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		Contacts:  constatcs,
		CreatedOn: time.Now(),
	}, nil
}
