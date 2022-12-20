package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type IActivitiesClientV1 interface {
	GetPartyActivities(ctx context.Context, correlationId string, filter *data.FilterParams,
		paging *data.PagingParams) (result data.DataPage[*PartyActivityV1], err error)

	LogPartyActivity(ctx context.Context, correlationId string, activity *PartyActivityV1) (result *PartyActivityV1, err error)

	BatchPartyActivities(ctx context.Context, correlationId string, activities []*PartyActivityV1) error

	DeletePartyActivities(ctx context.Context, correlationId string, filter *data.FilterParams) error
}
