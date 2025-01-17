package core

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	merrors "github.com/mishudark/errors"

	"github.com/oslokommune/okctl/pkg/jsonpatch"

	"github.com/oslokommune/okctl/pkg/client"
	"github.com/oslokommune/okctl/pkg/config/constant"
	"github.com/oslokommune/okctl/pkg/scaffold/resources"
	"github.com/spf13/afero"

	"sigs.k8s.io/yaml"
)

const (
	kustomizationFilename   = "kustomization.yaml"
	defaultManifestFilemode = 0o644 // u+rw g+r o+r
	defaultFolderMode       = 0o755 // a+rwx g-w o-w
)

type applicationManifestService struct {
	absoluteApplicationsDir string
	fs                      *afero.Afero
}

// SaveManifest saves a manifest to the application base directory and adds the manifest to the kustomization file
func (a *applicationManifestService) SaveManifest(_ context.Context, opts client.SaveManifestOpts) error {
	workingDir := path.Join(a.absoluteApplicationsDir, opts.ApplicationName, constant.DefaultApplicationBaseDir)

	err := a.fs.MkdirAll(workingDir, defaultFolderMode)
	if err != nil {
		return fmt.Errorf("creating application base directory: %w", err)
	}

	err = a.fs.WriteFile(path.Join(workingDir, opts.Filename), opts.Content, defaultManifestFilemode)
	if err != nil {
		return fmt.Errorf("writing manifest file: %w", err)
	}

	kustomizationManifest, err := acquireKustomizeFile(a.fs, workingDir)
	if err != nil {
		return fmt.Errorf("acquiring Kustomize file: %w", err)
	}

	kustomizationManifest.AddResource(opts.Filename)

	rawKustomizationManifest, err := yaml.Marshal(kustomizationManifest)
	if err != nil {
		return fmt.Errorf("marshalling kustomization manifest: %w", err)
	}

	err = a.fs.WriteFile(path.Join(workingDir, kustomizationFilename), rawKustomizationManifest, defaultManifestFilemode)
	if err != nil {
		return fmt.Errorf("writing to kustomization file: %w", err)
	}

	return nil
}

// SavePatch saves a json patch to the application overlay directory and adds the patch to the kustomization file
func (a *applicationManifestService) SavePatch(_ context.Context, opts client.SavePatchOpts) error {
	workingDir := path.Join(a.absoluteApplicationsDir, opts.ApplicationName, constant.DefaultApplicationOverlayDir, opts.ClusterName)
	patchFilename := fmt.Sprintf("%s-patch.json", strings.ToLower(opts.Kind))

	err := a.fs.MkdirAll(workingDir, defaultFolderMode)
	if err != nil {
		return fmt.Errorf("creating overlay directory: %w", err)
	}

	rawPatch, err := json.Marshal(opts.Patch)
	if err != nil {
		return fmt.Errorf("marshalling patch: %w", err)
	}

	err = a.fs.WriteFile(path.Join(workingDir, patchFilename), rawPatch, defaultManifestFilemode)
	if err != nil {
		return fmt.Errorf("writing patch file: %w", err)
	}

	kustomizationFile, err := acquireKustomizeFile(a.fs, workingDir)
	if err != nil {
		return fmt.Errorf("acquiring kustomization file: %w", err)
	}

	kustomizationFile.AddResource("../../base")

	kustomizationFile.AddPatch(resources.PatchReference{
		Path:   patchFilename,
		Target: resources.PatchTarget{Kind: opts.Kind},
	})

	rawKustomizationFile, err := yaml.Marshal(kustomizationFile)
	if err != nil {
		return fmt.Errorf("marshalling kustomization file: %w", err)
	}

	err = a.fs.WriteFile(path.Join(workingDir, kustomizationFilename), rawKustomizationFile, defaultManifestFilemode)
	if err != nil {
		return fmt.Errorf("writing kustomization file: %w", err)
	}

	return nil
}

// GetPatch retrieves a json patch from the application overlay directory
func (a *applicationManifestService) GetPatch(_ context.Context, opts client.GetPatchOpts) (jsonpatch.Patch, error) {
	workingDir := path.Join(a.absoluteApplicationsDir, opts.ApplicationName, constant.DefaultApplicationOverlayDir, opts.ClusterName)
	patchFilename := fmt.Sprintf("%s-patch.json", strings.ToLower(opts.Kind))

	content, err := a.fs.ReadFile(path.Join(workingDir, patchFilename))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return jsonpatch.Patch{}, merrors.E(err, "patch not found", merrors.NotExist)
		}

		return jsonpatch.Patch{}, fmt.Errorf("reading patch: %w", err)
	}

	patch := make([]jsonpatch.Operation, 0)

	err = json.Unmarshal(content, &patch)
	if err != nil {
		return jsonpatch.Patch{}, fmt.Errorf("unmarshalling patch: %w", err)
	}

	return jsonpatch.Patch{Operations: patch}, nil
}

func acquireKustomizeFile(fs *afero.Afero, absoluteDirPath string) (resources.Kustomization, error) {
	absoluteKustomizePath := path.Join(absoluteDirPath, kustomizationFilename)

	exists, err := fs.Exists(absoluteKustomizePath)
	if err != nil {
		return resources.Kustomization{}, fmt.Errorf("checking kustomization existence: %w", err)
	}

	if !exists {
		return *resources.NewKustomization(), nil
	}

	rawContent, err := fs.ReadFile(absoluteKustomizePath)
	if err != nil {
		return resources.Kustomization{}, fmt.Errorf("opening kustomization file: %w", err)
	}

	var manifest resources.Kustomization

	err = yaml.Unmarshal(rawContent, &manifest)
	if err != nil {
		return resources.Kustomization{}, fmt.Errorf("unmarshalling kustomization file: %w", err)
	}

	return manifest, nil
}

// NewApplicationManifestService initializes an Application Manifest Service
func NewApplicationManifestService(fs *afero.Afero, absoluteApplicationsDir string) client.ApplicationManifestService {
	return &applicationManifestService{
		absoluteApplicationsDir: absoluteApplicationsDir,
		fs:                      fs,
	}
}
