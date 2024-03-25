package campaign

import (
	"testing"

	"github.com/cristianmayco/go-emailn/internal/contract"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	service := Service{}
	newCampaign := contract.NewCampaignDTO{
		Name:    "test",
		Content: "body",
		Emails:  []string{"cristian@gmail.com", "cristianteste@gmail.com"},
	}

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_SaveCampaign(t *testing.T) {
	newCampaign := contract.NewCampaignDTO{
		Name:    "test",
		Content: "body",
		Emails:  []string{"cristian@gmail.com", "cristianteste@gmail.com"},
	}
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name ||
			campaign.Content != newCampaign.Content ||
			len(campaign.Contacts) != 2 ||
			campaign.Contacts[0].Email != newCampaign.Emails[0] ||
			campaign.Contacts[1].Email != newCampaign.Emails[1] {
			return false
		}
		return true

	})).Return(nil)
	service := Service{
		Repository: repositoryMock,
	}

	_, _ = service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)

}
