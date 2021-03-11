package main

import (
	"fmt"

	"github.com/oslokommune/okctl/pkg/apis/okctl.io/v1alpha1"
	"github.com/oslokommune/okctl/pkg/okctl"
	"github.com/spf13/cobra"
	"text/template"
)

const scaffoldClusterArgumentQuantity = 2

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
		Use:     "cluster CLUSTER_NAME ENVIRONMENT",
		Example: exampleUsage,
		Short:   "Scaffold cluster resource template",
		Long:    "Scaffolds a cluster resource which can be used to control cluster resources",
		Args:    cobra.ExactArgs(scaffoldClusterArgumentQuantity),
		PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]
			opts.Environment = args[1]

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
	flags.StringVarP(&opts.Organization, "github-organization", "o", "oslokommune", usageOrganization)
	flags.StringVarP(&opts.RepositoryName, "repository-name", "r", "my_iac_repo_name", usageRepository)
	flags.StringVarP(&opts.AWSAccountID, "aws-account-id", "i", "123456789123", usageAWSAccountID)

	return cmd
}

const (
	usageAWSAccountID = `the aws account where the resources provisioned by okctl should reside`
	usageOrganization = `the organization that owns the infrastructure-as-code repository`
	usageRepository   = `the name of the repository that will contain infrastructure-as-code`
	exampleUsage      = `okctl scaffold cluster utviklerportalen production > cluster.yaml`
)

const clusterTemplate = `apiVersion: okctl.io/v1alpha2
kind: Cluster

metadata:
  accountID: {{ .AWSAccountID }}
  environment: {{ .Environment }}
  name: {{ .Name }}
  region: eu-west-1

clusterRootURL: {{ .Name }}-{{ .Environment }}.oslo.systems

github:
  organisation: {{ .Organization }}
  outputPath: infrastructure
  repository: {{ .RepositoryName }}

integrations:
  argoCD: true
  autoscaler: true
  awsLoadBalancerController: true
  blockstorage: true
  cognito: true
  externalDNS: true
  externalSecrets: true
  kubePromStack: true

#vpc:
#  cidr: 192.168.0.0/20
#  highAvailability: true
`
