package v1alpha1_test

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"testing"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/yaml"

	"github.com/sebdah/goldie/v2"

	"github.com/oslokommune/okctl/pkg/apis/okctl.io/v1alpha1"
)

func TestCluster(t *testing.T) {
	testCases := []struct {
		name    string
		cluster v1alpha1.Cluster
		golden  string
	}{
		{
			name:    "Empty cluster",
			cluster: v1alpha1.Cluster{},
			golden:  "empty-cluster.yml",
		},
		{
			name: "Default cluster",
			cluster: v1alpha1.NewDefaultCluster(
				"okctl",
				"stage",
				"oslokommune",
				"okctl-iac",
				"123456789012",
			),
			golden: "default-cluster.yml",
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			got, err := yaml.Marshal(tc.cluster)
			assert.NoError(t, err)

			g := goldie.New(t)
			g.Assert(t, tc.golden, got)
		})
	}
}

func newPassingCluster() v1alpha1.Cluster {
	return v1alpha1.NewDefaultCluster(
		"x",
		"tre",
		"x",
		"x",
		"000000000000",
	)
}

func TestInvalidClusterValidations(t *testing.T) {
	testCases := []struct {
		name        string
		withCluster func() v1alpha1.Cluster
		expectError string
	}{
		{
			name: "Should pass when everything is A-ok",
			withCluster: func() v1alpha1.Cluster {
				return newPassingCluster()
			},
			expectError: "",
		},
		{
			name: "Should fail when name is empty",
			withCluster: func() v1alpha1.Cluster {
				c := newPassingCluster()
				
				c.Metadata.Name = ""

				return c
			},
			expectError: "",
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.withCluster().Validate()
			
			if tc.expectError != "" {
				assert.ErrorAs(t, err, testerr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
