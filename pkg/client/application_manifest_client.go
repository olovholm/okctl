package client

import (
	"context"

	"github.com/oslokommune/okctl/pkg/jsonpatch"
)

// SaveManifestOpts contains required data to save a Kubernetes manifest
type SaveManifestOpts struct {
	ApplicationName string

	Filename string
	Content  []byte
}

// SavePatchOpts contains required data to save a Kustomize patch
type SavePatchOpts struct {
	ApplicationName string
	ClusterName     string

	Kind  string
	Patch jsonpatch.Patch
}

// GetPatchOpts contains required data to retrieve a Kustomize patch
type GetPatchOpts struct {
	ApplicationName string
	ClusterName     string

	Kind string
}

// ApplicationManifestService defines functionality for the ApplicationManifestService
type ApplicationManifestService interface {
	// SaveManifest knows how to store a Kubernetes manifest and update Kustomize resources
	SaveManifest(ctx context.Context, opts SaveManifestOpts) error
	// SavePatch knows how to store a Kustomize patch
	SavePatch(ctx context.Context, opts SavePatchOpts) error
	// GetPatch knows how to retrieve a Kustomize patch
	GetPatch(ctx context.Context, opts GetPatchOpts) (jsonpatch.Patch, error)
}
