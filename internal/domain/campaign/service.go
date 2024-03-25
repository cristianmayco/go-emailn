package campaign

import (
	"github.com/cristianmayco/go-emailn/internal/contract"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaing contract.NewCampaignDTO) (string, error) {

	campaign, _ := NewCampaign(newCampaing.Name, newCampaing.Content, newCampaing.Emails)

	s.Repository.Save(campaign)

	return campaign.ID, nil
}
