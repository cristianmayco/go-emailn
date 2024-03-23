package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t)
	name := "test"
	content := "body"
	contacts := []string{"cristian@gmail.com", "cristianteste@gmail.com"}

	campaign := NewCampaign(name, content, contacts)

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), 2)
	assert.Equal(campaign.Contacts[0].Email, "cristian@gmail.com")
	assert.Equal(campaign.Contacts[1].Email, "cristianteste@gmail.com")

}

func Test_NewCampaign_IDIsNotNil(t *testing.T) {
	assert := assert.New(t)
	name := "test"
	content := "body"
	contacts := []string{"cristian@gmail.com", "cristianteste@gmail.com"}

	campaign := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreatedOnIsNotNil(t *testing.T) {
	assert := assert.New(t)
	name := "test"
	content := "body"
	contacts := []string{"cristian@gmail.com", "cristianteste@gmail.com"}

	campaign := NewCampaign(name, content, contacts)

	assert.Greater(time.Now(), campaign.CreatedOn)
}
