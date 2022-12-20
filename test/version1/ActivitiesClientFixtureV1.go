package test_version1

import (
	"context"
	"testing"
	"time"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"

	"github.com/pip-services-users2/client-activities-go/version1"
	"github.com/stretchr/testify/assert"
)

type ActivitiesClientFixtureV1 struct {
	Client version1.IActivitiesClientV1
}

var ACTIVITY = &version1.PartyActivityV1{
	Id:   "",
	Type: "test",
	Time: time.Now(),
	Party: &version1.ReferenceV1{
		Id:   "1",
		Type: "party",
		Name: "Test User",
	},
	RefItem: &version1.ReferenceV1{
		Id:   "2",
		Type: "party",
		Name: "Admin User",
	},
	RefParents: []*version1.ReferenceV1{},
	RefParty:   nil,
	Details:    nil,
}

func NewActivitiesClientFixtureV1(client version1.IActivitiesClientV1) *ActivitiesClientFixtureV1 {
	return &ActivitiesClientFixtureV1{
		Client: client,
	}
}

func (c *ActivitiesClientFixtureV1) clear() {
	c.Client.DeletePartyActivities(context.Background(), "", nil)
}

func (c *ActivitiesClientFixtureV1) TestBatchPartyActivities(t *testing.T) {
	c.clear()
	defer c.clear()

	// Log an activity batch
	err := c.Client.BatchPartyActivities(context.Background(), "", []*version1.PartyActivityV1{ACTIVITY, ACTIVITY, ACTIVITY})
	assert.Nil(t, err)

	// Get activities
	page, err1 := c.Client.GetPartyActivities(context.Background(), "", data.NewFilterParamsFromTuples("party_id", "1"), nil)
	assert.Nil(t, err1)

	assert.NotNil(t, page)
	assert.True(t, len(page.Data) > 2)

	activity := page.Data[0]

	assert.NotNil(t, activity.Time)
	assert.Equal(t, activity.Type, ACTIVITY.Type)
	assert.Equal(t, activity.Party.Name, ACTIVITY.Party.Name)
}
