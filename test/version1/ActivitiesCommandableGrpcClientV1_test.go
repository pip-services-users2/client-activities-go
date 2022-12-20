package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-users2/client-activities-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type applicationsCommandableGrpcClientV1Test struct {
	client  *version1.ActivitiesCommandableGrpcClientV1
	fixture *ActivitiesClientFixtureV1
}

func newApplicationsCommandableGrpcClientV1Test() *applicationsCommandableGrpcClientV1Test {
	return &applicationsCommandableGrpcClientV1Test{}
}

func (c *applicationsCommandableGrpcClientV1Test) setup(t *testing.T) *ActivitiesClientFixtureV1 {
	var GRPC_HOST = os.Getenv("GRPC_HOST")
	if GRPC_HOST == "" {
		GRPC_HOST = "localhost"
	}
	var GRPC_PORT = os.Getenv("GRPC_PORT")
	if GRPC_PORT == "" {
		GRPC_PORT = "8090"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", GRPC_HOST,
		"connection.port", GRPC_PORT,
	)

	c.client = version1.NewActivitiesCommandableGrpcClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewActivitiesClientFixtureV1(c.client)

	return c.fixture
}

func (c *applicationsCommandableGrpcClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestCommandableGrpcBatchPartyActivities(t *testing.T) {
	c := newApplicationsCommandableGrpcClientV1Test()
	c.setup(t)
	defer c.teardown(t)

	c.fixture.TestBatchPartyActivities(t)
}
