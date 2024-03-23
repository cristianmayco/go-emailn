package campaign

import "time"

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

func NewCampaign(name string, content string, emails []string) *Campaign {
	constatcs := make([]Contatc, len(emails))

	for index, email := range emails {
		constatcs[index].Email = email
	}

	return &Campaign{
		Name:      name,
		Content:   content,
		Contacts:  constatcs,
		CreatedOn: time.Now(),
	}
}
