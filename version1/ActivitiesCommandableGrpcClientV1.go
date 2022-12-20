package version1

import (
	"context"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-grpc-gox/clients"
)

type ActivitiesCommandableGrpcClientV1 struct {
	*clients.CommandableGrpcClient
}

func NewActivitiesCommandableGrpcClientV1() *ActivitiesCommandableGrpcClientV1 {
	return NewActivitiesCommandableGrpcClientV1WithConfig(nil)
}

func NewActivitiesCommandableGrpcClientV1WithConfig(config *cconf.ConfigParams) *ActivitiesCommandableGrpcClientV1 {
	c := &ActivitiesCommandableGrpcClientV1{
		CommandableGrpcClient: clients.NewCommandableGrpcClient("v1/activities"),
	}

	if config != nil {
		c.Configure(context.Background(), config)
	}

	return c
}

func (c *ActivitiesCommandableGrpcClientV1) GetPartyActivities(ctx context.Context, correlationId string, filter *cdata.FilterParams,
	paging *cdata.PagingParams) (result cdata.DataPage[*PartyActivityV1], err error) {
	res, err := c.CallCommand(ctx, "get_party_activities", correlationId, cdata.NewAnyValueMapFromTuples(
		"filter", filter,
		"paging", paging,
	))

	if err != nil {
		return *cdata.NewEmptyDataPage[*PartyActivityV1](), err
	}

	return clients.HandleHttpResponse[cdata.DataPage[*PartyActivityV1]](res, correlationId)
}

func (c *ActivitiesCommandableGrpcClientV1) LogPartyActivity(ctx context.Context, correlationId string, activity *PartyActivityV1) (result *PartyActivityV1, err error) {
	res, err := c.CallCommand(ctx, "log_party_activity", correlationId, cdata.NewAnyValueMapFromTuples(
		"activity", activity,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*PartyActivityV1](res, correlationId)
}

func (c *ActivitiesCommandableGrpcClientV1) BatchPartyActivities(ctx context.Context, correlationId string, activities []*PartyActivityV1) error {
	_, err := c.CallCommand(ctx, "batch_party_activities", correlationId, cdata.NewAnyValueMapFromTuples(
		"activities", activities,
	))

	return err
}

func (c *ActivitiesCommandableGrpcClientV1) DeletePartyActivities(ctx context.Context, correlationId string, filter *cdata.FilterParams) error {
	_, err := c.CallCommand(ctx, "delete_party_activities", correlationId, cdata.NewAnyValueMapFromTuples(
		"filter", filter,
	))

	return err
}
