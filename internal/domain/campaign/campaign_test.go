package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCampaign(t *testing.T) {
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
