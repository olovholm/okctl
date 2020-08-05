// Package mock provides mocks for use with tests
package mock

import (
	"context"
	"fmt"

	"github.com/oslokommune/okctl/pkg/api"
	"github.com/oslokommune/okctl/pkg/api/okctl.io/v1alpha1"
)

const (
	// DefaultEnv is a default environment used in mocks
	DefaultEnv = "pro"
	// DefaultAWSAccountID is a default aws account id used in mocks
	DefaultAWSAccountID = "123456789012"
	// DefaultCidr is a default cidr used in mocks
	DefaultCidr = "192.168.0.0/20"
	// DefaultRegion is a default aws region used in mocks
	DefaultRegion = "eu-west-1"
	// DefaultAvailabilityZone is a default aws availability zone used in mocks
	DefaultAvailabilityZone = "eu-west-1a"
	// DefaultRepositoryName is a default git repo name used in mocks
	DefaultRepositoryName = "test"
	// DefaultClusterName is a default eks cluster name used in mocks
	DefaultClusterName = "test-cluster-pro"
	// DefaultVpcID is a default aws vpc id used in mocks
	DefaultVpcID = "vpc-0e9801d129EXAMPLE"
	// DefaultPublicSubnetID is a default aws public subnet id used in mocks
	DefaultPublicSubnetID = "subnet-0bb1c79de3EXAMPLE"
	// DefaultPublicSubnetCidr is a default public subnet cidr used in mocks
	DefaultPublicSubnetCidr = "192.168.1.0/24"
	// DefaultPrivateSubnetID is a default private aws subnet id used in mocks
	DefaultPrivateSubnetID = "subnet-8EXAMPLE"
	// DefaultPrivateSubnetCidr is a default private subnet cidr used in mocks
	DefaultPrivateSubnetCidr = "192.168.2.0/24"
	// DefaultVpcStackName is the default stack name for a vpc
	DefaultVpcStackName = "test-vpc-pro"
	// DefaultVpcCloudFormationTemplate is a default cloud formation template
	DefaultVpcCloudFormationTemplate = "something"
)

// ErrBad just defines a mocked error
var ErrBad = fmt.Errorf("something bad")

// DefaultVpcCreateOpts returns options for creating a vpc with defaults set
func DefaultVpcCreateOpts() api.CreateVpcOpts {
	return api.CreateVpcOpts{
		AwsAccountID: DefaultAWSAccountID,
		ClusterName:  DefaultClusterName,
		Env:          DefaultEnv,
		RepoName:     DefaultRepositoryName,
		Cidr:         DefaultCidr,
		Region:       DefaultRegion,
	}
}

// DefaultVpcDeleteOpts returns options for deleting a vpc with defaults set
func DefaultVpcDeleteOpts() api.DeleteVpcOpts {
	return api.DeleteVpcOpts{
		Env:      DefaultEnv,
		RepoName: DefaultRepositoryName,
	}
}

// DefaultClusterDeleteOpts returns options for deleting a cluster with defaults set
func DefaultClusterDeleteOpts() api.ClusterDeleteOpts {
	return api.ClusterDeleteOpts{
		Environment:    DefaultEnv,
		RepositoryName: DefaultRepositoryName,
	}
}

// DefaultClusterCreateOpts returns options for creating a cluster with defaults set
func DefaultClusterCreateOpts() api.ClusterCreateOpts {
	return api.ClusterCreateOpts{
		Environment:    DefaultEnv,
		AWSAccountID:   DefaultAWSAccountID,
		Cidr:           DefaultCidr,
		RepositoryName: DefaultRepositoryName,
		Region:         DefaultRegion,
		ClusterName:    DefaultClusterName,
	}
}

// DefaultCreateClusterConfigOpts returns options for creating a cluster configuration
// with defaults set
func DefaultCreateClusterConfigOpts() api.CreateClusterConfigOpts {
	return api.CreateClusterConfigOpts{
		ClusterName:  DefaultClusterName,
		Region:       DefaultRegion,
		Cidr:         DefaultCidr,
		AwsAccountID: DefaultAWSAccountID,
	}
}

// DefaultPublicSubnets returns a map of public subnets with defaults set
func DefaultPublicSubnets() map[string]api.ClusterNetwork {
	return map[string]api.ClusterNetwork{
		DefaultAvailabilityZone: {
			ID:   DefaultPublicSubnetID,
			CIDR: DefaultPublicSubnetCidr,
		},
	}
}

// DefaultVpcPublicSubnets returns a list of public subnets with defaults set
func DefaultVpcPublicSubnets() []api.VpcSubnet {
	return []api.VpcSubnet{
		{
			ID:               DefaultPublicSubnetID,
			Cidr:             DefaultPublicSubnetCidr,
			AvailabilityZone: DefaultAvailabilityZone,
		},
	}
}

// DefaultPrivateSubnets returns a map of private subnets with defaults set
func DefaultPrivateSubnets() map[string]api.ClusterNetwork {
	return map[string]api.ClusterNetwork{
		DefaultAvailabilityZone: {
			ID:   DefaultPrivateSubnetID,
			CIDR: DefaultPrivateSubnetCidr,
		},
	}
}

// DefaultVpcPrivateSubnets returns a list of private subnets with defaults set
func DefaultVpcPrivateSubnets() []api.VpcSubnet {
	return []api.VpcSubnet{
		{
			ID:               DefaultPrivateSubnetID,
			Cidr:             DefaultPrivateSubnetCidr,
			AvailabilityZone: DefaultAvailabilityZone,
		},
	}
}

// DefaultClusterConfig returns a cluster config with defaults set
func DefaultClusterConfig() *api.ClusterConfig {
	cfg := api.NewClusterConfig()

	cfg.Metadata.Name = DefaultClusterName
	cfg.Metadata.Region = DefaultRegion

	cfg.VPC.ID = DefaultVpcID
	cfg.VPC.CIDR = DefaultCidr

	cfg.VPC.Subnets.Public = DefaultPublicSubnets()
	cfg.VPC.Subnets.Private = DefaultPrivateSubnets()

	cfg.IAM.FargatePodExecutionRolePermissionsBoundary = v1alpha1.PermissionsBoundaryARN(DefaultAWSAccountID)
	cfg.IAM.ServiceRolePermissionsBoundary = v1alpha1.PermissionsBoundaryARN(DefaultAWSAccountID)

	return cfg
}

// DefaultVpc returns a vpc with defaults set
func DefaultVpc() *api.Vpc {
	return &api.Vpc{
		StackName:              DefaultVpcStackName,
		CloudFormationTemplate: []byte(DefaultVpcCloudFormationTemplate),
		ID:                     DefaultVpcID,
		PublicSubnets:          DefaultVpcPublicSubnets(),
		PrivateSubnets:         DefaultVpcPrivateSubnets(),
	}
}

// DefaultCluster returns an api cluster definition with defaults set
func DefaultCluster() *api.Cluster {
	return &api.Cluster{
		Environment:  DefaultEnv,
		AWSAccountID: DefaultAWSAccountID,
		Cidr:         DefaultCidr,
		Config:       DefaultClusterConfig(),
	}
}

// ClusterService provides a mock for the cluster service interface
type ClusterService struct {
	CreateClusterFn func(context.Context, api.ClusterCreateOpts) (*api.Cluster, error)
	DeleteClusterFn func(context.Context, api.ClusterDeleteOpts) error
}

// CreateCluster invokes a mocked function to create a cluster
func (s *ClusterService) CreateCluster(ctx context.Context, opts api.ClusterCreateOpts) (*api.Cluster, error) {
	return s.CreateClusterFn(ctx, opts)
}

// DeleteCluster invokes a mocked function to delete a cluster
func (s *ClusterService) DeleteCluster(ctx context.Context, opts api.ClusterDeleteOpts) error {
	return s.DeleteClusterFn(ctx, opts)
}

// NewGoodClusterService returns a cluster service that will succeed
func NewGoodClusterService() *ClusterService {
	return &ClusterService{
		CreateClusterFn: func(_ context.Context, opts api.ClusterCreateOpts) (*api.Cluster, error) {
			return DefaultCluster(), nil
		},
		DeleteClusterFn: func(context.Context, api.ClusterDeleteOpts) error {
			return nil
		},
	}
}

// NewBadClusterService returns a cluster service that will fail
func NewBadClusterService() *ClusterService {
	return &ClusterService{
		CreateClusterFn: func(context.Context, api.ClusterCreateOpts) (*api.Cluster, error) {
			return nil, ErrBad
		},
		DeleteClusterFn: func(context.Context, api.ClusterDeleteOpts) error {
			return ErrBad
		},
	}
}

// VpcCloud provides a mock for the cluster cloud interface
type VpcCloud struct {
	CreateVpcFn func(opts api.CreateVpcOpts) (*api.Vpc, error)
	DeleteVpcFn func(opts api.DeleteVpcOpts) error
}

// CreateVpc invokes the mocked create cluster function
func (c *VpcCloud) CreateVpc(opts api.CreateVpcOpts) (*api.Vpc, error) {
	return c.CreateVpcFn(opts)
}

// DeleteVpc invokes the mocked delete cluster function
func (c *VpcCloud) DeleteVpc(opts api.DeleteVpcOpts) error {
	return c.DeleteVpcFn(opts)
}

// NewGoodVpcCloud returns a cluster cloud that will succeed
func NewGoodVpcCloud() *VpcCloud {
	return &VpcCloud{
		CreateVpcFn: func(opts api.CreateVpcOpts) (*api.Vpc, error) {
			return DefaultVpc(), nil
		},
		DeleteVpcFn: func(opts api.DeleteVpcOpts) error {
			return nil
		},
	}
}

// NewBadVpcCloud returns a cluster cloud that will fail
func NewBadVpcCloud() *VpcCloud {
	return &VpcCloud{
		CreateVpcFn: func(opts api.CreateVpcOpts) (*api.Vpc, error) {
			return nil, ErrBad
		},
		DeleteVpcFn: func(opts api.DeleteVpcOpts) error {
			return ErrBad
		},
	}
}

// ClusterExe provides a mock for the cluster exe interface
type ClusterExe struct {
	CreateClusterFn func(*api.ClusterConfig) error
	DeleteClusterFn func(*api.ClusterConfig) error
}

// CreateCluster invokes the mocked create cluster function
func (c *ClusterExe) CreateCluster(config *api.ClusterConfig) error {
	return c.CreateClusterFn(config)
}

// DeleteCluster invokes the mocked delete cluster function
func (c *ClusterExe) DeleteCluster(config *api.ClusterConfig) error {
	return c.DeleteClusterFn(config)
}

// NewGoodClusterExe returns a cluster exe that will succeed
func NewGoodClusterExe() *ClusterExe {
	return &ClusterExe{
		CreateClusterFn: func(config *api.ClusterConfig) error {
			return nil
		},
		DeleteClusterFn: func(config *api.ClusterConfig) error {
			return nil
		},
	}
}

// NewBadClusterExe returns a cluster exe that will fail
func NewBadClusterExe() *ClusterExe {
	return &ClusterExe{
		CreateClusterFn: func(config *api.ClusterConfig) error {
			return ErrBad
		},
		DeleteClusterFn: func(config *api.ClusterConfig) error {
			return ErrBad
		},
	}
}

// ClusterStore provides a mock for the cluster store interface
type ClusterStore struct {
	SaveClusterFn   func(*api.Cluster) error
	DeleteClusterFn func(env string) error
	GetClusterFn    func(env string) (*api.Cluster, error)
}

// SaveCluster invokes the mocked save function
func (c *ClusterStore) SaveCluster(cluster *api.Cluster) error {
	return c.SaveClusterFn(cluster)
}

// DeleteCluster invokes the mocked delete function
func (c *ClusterStore) DeleteCluster(env string) error {
	return c.DeleteClusterFn(env)
}

// GetCluster invokes the mocked get function
func (c *ClusterStore) GetCluster(env string) (*api.Cluster, error) {
	return c.GetClusterFn(env)
}

// NewGoodClusterStore returns a cluster store that will succeed
func NewGoodClusterStore() *ClusterStore {
	return &ClusterStore{
		SaveClusterFn: func(cluster *api.Cluster) error {
			return nil
		},
		DeleteClusterFn: func(env string) error {
			return nil
		},
		GetClusterFn: func(env string) (*api.Cluster, error) {
			return DefaultCluster(), nil
		},
	}
}

// NewBadClusterStore returns a cluster store that will fail
func NewBadClusterStore() *ClusterStore {
	return &ClusterStore{
		SaveClusterFn: func(cluster *api.Cluster) error {
			return ErrBad
		},
		DeleteClusterFn: func(env string) error {
			return ErrBad
		},
		GetClusterFn: func(env string) (*api.Cluster, error) {
			return nil, ErrBad
		},
	}
}

// ClusterConfigStore provides a mock for the cluster config store
type ClusterConfigStore struct {
	SaveClusterConfigFn   func(*api.ClusterConfig) error
	DeleteClusterConfigFn func(env string) error
	GetClusterConfigFn    func(env string) (*api.ClusterConfig, error)
}

// SaveClusterConfig invokes the mocked save cluster config function
func (c *ClusterConfigStore) SaveClusterConfig(config *api.ClusterConfig) error {
	return c.SaveClusterConfigFn(config)
}

// DeleteClusterConfig invokes the mocked delete cluster config function
func (c *ClusterConfigStore) DeleteClusterConfig(env string) error {
	return c.DeleteClusterConfigFn(env)
}

// GetClusterConfig invokes the mocked get cluster config function
func (c *ClusterConfigStore) GetClusterConfig(env string) (*api.ClusterConfig, error) {
	return c.GetClusterConfigFn(env)
}

// NewGoodClusterConfigStore returns a cluster config store that will succeed
func NewGoodClusterConfigStore() *ClusterConfigStore {
	return &ClusterConfigStore{
		SaveClusterConfigFn: func(config *api.ClusterConfig) error {
			return nil
		},
		DeleteClusterConfigFn: func(env string) error {
			return nil
		},
		GetClusterConfigFn: func(env string) (*api.ClusterConfig, error) {
			return DefaultClusterConfig(), nil
		},
	}
}