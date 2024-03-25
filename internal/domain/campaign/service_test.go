package campaign

import (
	"errors"
	"testing"

	"github.com/cristianmayco/go-emailn/internal/contract"
	internalerros "github.com/cristianmayco/go-emailn/internal/internal_erros"
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

var (
	newCampaign = contract.NewCampaignDTO{
		Name:    "test",
		Content: "body",
		Emails:  []string{"cristian@gmail.com", "cristianteste@gmail.com"},
	}
	service = Service{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	service := Service{}

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_SaveCampaign(t *testing.T) {
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

	service.Repository = repositoryMock
	_, _ = service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)

}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)
	newCampaign.Name = ""

	_, err := service.Create(newCampaign)

	assert.NotNil(err)
	assert.Equal("name is required", err.Error())
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save on database"))

	service.Repository = repositoryMock
	_, err := service.Create(newCampaign)

	assert.True(errors.Is(internalerros.ErrInternal, err))

}
