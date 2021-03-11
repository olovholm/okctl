package main

import (
	"fmt"
	"text/template"

	"github.com/oslokommune/okctl/pkg/okctl"
	"github.com/spf13/cobra"
)

const scaffoldClusterArgumentQuantity = 0

type scaffoldClusterOpts struct {
	Name string

	AWSAccountID   string
	Environment    string
	Organization   string
	RepositoryName string
}

func buildScaffoldClusterCommand(o *okctl.Okctl) *cobra.Command {
	opts := scaffoldClusterOpts{}

	cmd := &cobra.Command{
		Use:     "cluster",
		Example: exampleUsage,
		Short:   "Scaffold cluster resource template",
		Long:    "Scaffolds a cluster resource which can be used to control cluster resources",
		Args:    cobra.ExactArgs(scaffoldClusterArgumentQuantity),
		PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			t, err := template.New("cluster.yaml").Parse(clusterTemplate)
			if err != nil {
				return fmt.Errorf("parsing template string: %w", err)
			}

			err = t.Execute(o.Out, opts)
			if err != nil {
				return fmt.Errorf("interpolating template: %w", err)
			}

			return nil
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.Name, "name", "n", "my-product-name", usageName)
	flags.StringVarP(&opts.Environment, "environment", "e", "development", usageEnvironment)
	flags.StringVarP(&opts.Organization, "github-organization", "o", "oslokommune", usageOrganization)
	flags.StringVarP(&opts.RepositoryName, "repository-name", "r", "my_iac_repo_name", usageRepository)
	flags.StringVarP(&opts.AWSAccountID, "aws-account-id", "i", "123456789123", usageAWSAccountID)

	return cmd
}

const (
	usageName         = `the name of the cluster`
	usageEnvironment  = `the environment for the cluster, for example dev or production`
	usageAWSAccountID = `the aws account where the resources provisioned by okctl should reside`
	usageOrganization = `the organization that owns the infrastructure-as-code repository`
	usageRepository   = `the name of the repository that will contain infrastructure-as-code`
	exampleUsage      = `okctl scaffold cluster utviklerportalen production > cluster.yaml`
)

const clusterTemplate = `apiVersion: okctl.io/v1alpha2
kind: Cluster

# For help finding values, see https://okctl.io/getting-started/create-cluster
metadata:
  # Account ID is your AWS account ID
  accountID: {{ .AWSAccountID }}
  # Environment is the name you use to identify the type of cluster it is. Common names are production, test, staging
  environment: {{ .Environment }}
  # Name can be anything, but should define the scope of the cluster. Meaning if the cluster is scoped to one product,
  # you might want to name it the name of the product. If the cluster contains all services and products owned by a
  # team, the team name might be more fitting.
  name: {{ .Name }}
  # Region defines the AWS region to prefer when creating resources
  region: eu-west-1

# The cluster root URL defines the domain of which to create services beneath. For example; okctl will setup ArgoCD
# which has a frontend. The frontend will be available at https://argocd.<clusterRootURL>. For Cognito it will be 
# https://auth.<clusterRootURL>
clusterRootURL: {{ .Name }}-{{ .Environment }}.oslo.systems

# For okctl to be able to setup ArgoCD correctly for you, it needs to know what repository on Github that will contain
# your infrastructure.
github:
  # The organization that owns the repository
  organisation: {{ .Organization }}
  # The folder to place infrastructure declarations
  outputPath: infrastructure
  # The name of the repository
  repository: {{ .RepositoryName }}

integrations:
  # ArgoCD is a service that watches a repository for Kubernetes charts and ensures the defined resources are running
  # as declared in the cluster
  argoCD: true
  # Autoscaler automatically adjusts the size of pods and nodes in your cluster depending on load
  autoscaler: true
  # AWS Load Balancer Controller handles routing from the internet to your application running inside your okctl
  # Kubernetes cluster. If you want your applications and services accessible from the internet, this needs to be
  # enabled.
  awsLoadBalancerController: true
  # Block storage provides persistent storage for your cluster (Persistent Volumes)
  blockstorage: true
  # Cognito is an authentication provider that okctl uses to control access to different resources, like ArgoCD and
  # Grafana
  cognito: true
  # External DNS handles defining the necessary DNS records required to route traffic to your defined service or 
  # application
  externalDNS: true
  # External Secrets automatically AWS secrets  
  externalSecrets: true
  kubePromStack: true

# okctl creates a Virtual Private Cloud for you which it organizes all the intended resources that require networking.
#vpc:
  # CIDR defines the VPC IP range. Leave this be if you don't know what it is/does
#  cidr: 192.168.0.0/20
#  highAvailability: true
`
