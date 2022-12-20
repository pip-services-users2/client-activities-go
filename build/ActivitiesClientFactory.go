package build

import (
	clients1 "github.com/pip-services-users2/client-activities-go/version1"

	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

type EmailSettingsClientFactory struct {
	*cbuild.Factory
}

func NewEmailSettingsClientFactory() *EmailSettingsClientFactory {
	c := &EmailSettingsClientFactory{
		Factory: cbuild.NewFactory(),
	}

	nullClientDescriptor := cref.NewDescriptor("service-activities", "client", "null", "*", "1.0")
	cmdHttpClientDescriptor := cref.NewDescriptor("service-activities", "client", "commandable-http", "*", "1.0")
	cmdGrpcClientDescriptor := cref.NewDescriptor("service-activities", "client", "commandable-grpc", "*", "1.0")
	grpcClientDescriptor := cref.NewDescriptor("service-activities", "client", "grpc", "*", "1.0")

	c.RegisterType(nullClientDescriptor, clients1.NewActivitiesNullClientV1)
	c.RegisterType(cmdHttpClientDescriptor, clients1.NewActivitiesCommandableHttpClientV1)
	c.RegisterType(cmdGrpcClientDescriptor, clients1.NewActivitiesCommandableGrpcClientV1)
	c.RegisterType(grpcClientDescriptor, clients1.NewActivitiesGrpcClientV1)

	return c
}
