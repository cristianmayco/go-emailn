package campaign

import (
	"time"

	"github.com/rs/xid"
)

type Contatc struct {
	Email string `validate:"email"`
}

type Campaign struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"min=3,max=24"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"min=5,max=1024"`
	Contacts  []Contatc `validate:"min=1,dive"`
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

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
