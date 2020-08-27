package core

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/oslokommune/okctl/pkg/api"
)

func makeCreateExternalDNSKubeDeploymentEndpoint(s api.KubeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return s.CreateExternalDNSKubeDeployment(ctx, request.(api.CreateExternalDNSKubeDeploymentOpts))
	}
}