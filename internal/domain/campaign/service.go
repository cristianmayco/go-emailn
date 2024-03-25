package campaign

import (
	"github.com/cristianmayco/go-emailn/internal/contract"
	internalerros "github.com/cristianmayco/go-emailn/internal/internal_erros"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaing contract.NewCampaignDTO) (string, error) {

	campaign, err := NewCampaign(newCampaing.Name, newCampaing.Content, newCampaing.Emails)

	if err != nil {
		return "", err
	}

	err = s.Repository.Save(campaign)

	if err != nil {
		return "", internalerros.ErrInternal

	}

	return campaign.ID, nil
}
