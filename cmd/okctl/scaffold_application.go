package main

import (
	"fmt"

	"github.com/oslokommune/okctl/pkg/commands"

	"github.com/oslokommune/okctl/pkg/okctl"
	"github.com/spf13/cobra"
)

const requiredArgumentsForCreateApplicationCommand = 0

// nolint: funlen
func buildScaffoldApplicationCommand(o *okctl.Okctl) *cobra.Command {
	opts := commands.ScaffoldApplicationOpts{}

	cmd := &cobra.Command{
		Use:   "application",
		Short: ScaffoldShortDescription,
		Long:  ScaffoldLongDescription,
		Args:  cobra.ExactArgs(requiredArgumentsForCreateApplicationCommand),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if declarationPath != "" {
				clusterDeclaration, err := commands.InferClusterFromStdinOrFile(o.In, declarationPath)
				if err != nil {
					return fmt.Errorf("inferring cluster declaration: %w", err)
				}

				opts.PrimaryHostedZone = clusterDeclaration.ClusterRootDomain
			} else {
				opts.PrimaryHostedZone = "okctl.io"
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return commands.ScaffoldApplicationDeclaration(o.Out, opts)
		},
	}

	return cmd
}
