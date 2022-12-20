package version1

import (
	"context"

	"github.com/pip-services-users2/client-activities-go/protos"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-grpc-gox/clients"
)

type ActivitiesGrpcClientV1 struct {
	*clients.GrpcClient
}

func NewActivitiesGrpcClientV1() *ActivitiesGrpcClientV1 {
	return &ActivitiesGrpcClientV1{
		GrpcClient: clients.NewGrpcClient("activities_v1.Activities"),
	}
}

func (c *ActivitiesGrpcClientV1) GetPartyActivities(ctx context.Context, correlationId string, filter *data.FilterParams,
	paging *data.PagingParams) (result data.DataPage[*PartyActivityV1], err error) {
	timing := c.Instrument(ctx, correlationId, "activities_v1.get_party_activities")
	defer timing.EndTiming(ctx, err)

	req := &protos.PartyActivityPageRequest{
		CorrelationId: correlationId,
	}
	if filter != nil {
		req.Filter = filter.Value()
	}
	if paging != nil {
		req.Paging = &protos.PagingParams{
			Skip:  paging.GetSkip(0),
			Take:  (int32)(paging.GetTake(100)),
			Total: paging.Total,
		}
	}

	reply := new(protos.PartyActivityPageReply)
	err = c.CallWithContext(ctx, "get_party_activities", correlationId, req, reply)
	if err != nil {
		return *data.NewEmptyDataPage[*PartyActivityV1](), err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return *data.NewEmptyDataPage[*PartyActivityV1](), err
	}

	result = toPartyActivityPage(reply.Page)

	return result, nil
}

func (c *ActivitiesGrpcClientV1) LogPartyActivity(ctx context.Context, correlationId string, activity *PartyActivityV1) (result *PartyActivityV1, err error) {
	timing := c.Instrument(ctx, correlationId, "activities_v1.log_party_activity")
	defer timing.EndTiming(ctx, err)

	req := &protos.PartyActivityLogRequest{
		CorrelationId: correlationId,
		Activity:      fromPartyActivity(activity),
	}

	reply := new(protos.PartyActivityObjectReply)
	err = c.CallWithContext(ctx, "log_party_activity", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toPartyActivity(reply.Activity)

	return result, nil
}

func (c *ActivitiesGrpcClientV1) BatchPartyActivities(ctx context.Context, correlationId string, activities []*PartyActivityV1) (err error) {
	timing := c.Instrument(ctx, correlationId, "activities_v1.batch_party_activities")
	defer timing.EndTiming(ctx, err)

	req := &protos.PartyActivityBatchRequest{
		CorrelationId: correlationId,
		Activities:    fromPartyActivities(activities),
	}

	reply := new(protos.PartyActivityOnlyErrorReply)
	err = c.CallWithContext(ctx, "batch_party_activities", correlationId, req, reply)
	if err != nil {
		return err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return err
	}

	return nil
}

func (c *ActivitiesGrpcClientV1) DeletePartyActivities(ctx context.Context, correlationId string, filter *data.FilterParams) (err error) {
	timing := c.Instrument(ctx, correlationId, "activities_v1.delete_party_activities")
	defer timing.EndTiming(ctx, err)

	req := &protos.PartyActivityDeleteRequest{
		CorrelationId: correlationId,
	}
	if filter != nil {
		req.Filter = filter.Value()
	}

	reply := new(protos.PartyActivityOnlyErrorReply)
	err = c.CallWithContext(ctx, "delete_party_activities", correlationId, req, reply)
	if err != nil {
		return err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return err
	}

	return nil
}
