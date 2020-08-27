package core

import (
	"context"

	"github.com/mishudark/errors"
	"github.com/oslokommune/okctl/pkg/api"
)

type kubeService struct {
	run   api.KubeRun
	store api.KubeStore
}

func (k *kubeService) CreateExternalDNSKubeDeployment(_ context.Context, opts api.CreateExternalDNSKubeDeploymentOpts) (*api.Kube, error) {
	err := opts.Validate()
	if err != nil {
		return nil, errors.E(err, "failed to validate input options")
	}

	kube, err := k.run.CreateExternalDNSKubeDeployment(opts)
	if err != nil {
		return nil, errors.E(err, "failed to deploy kubernetes manifests")
	}

	err = k.store.SaveExternalDNSKubeDeployment(kube)
	if err != nil {
		return nil, errors.E(err, "failed to save kubernetes manifests")
	}

	return kube, nil
}

// NewKubeService returns an initialised kube service
func NewKubeService(store api.KubeStore, run api.KubeRun) api.KubeService {
	return &kubeService{
		run:   run,
		store: store,
	}
}