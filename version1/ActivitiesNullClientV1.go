package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type ActivitiesNullClientV1 struct {
}

func NewActivitiesNullClientV1() *ActivitiesNullClientV1 {
	return &ActivitiesNullClientV1{}
}

func (c *ActivitiesNullClientV1) GetPartyActivities(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (result data.DataPage[*PartyActivityV1], err error) {
	return *data.NewEmptyDataPage[*PartyActivityV1](), nil
}

func (c *ActivitiesNullClientV1) LogPartyActivity(ctx context.Context, correlationId string, activity *PartyActivityV1) (result *PartyActivityV1, err error) {
	return nil, nil
}

func (c *ActivitiesNullClientV1) BatchPartyActivities(ctx context.Context, correlationId string, activities []*PartyActivityV1) error {
	return nil
}

func (c *ActivitiesNullClientV1) DeletePartyActivities(ctx context.Context, correlationId string, filter *data.FilterParams) error {
	return nil
}
