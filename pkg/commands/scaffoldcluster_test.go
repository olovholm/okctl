package commands

import (
	"bytes"
	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClusterDeclarationScaffold(t *testing.T) {
	testCases := []struct {
		name string

		withOpts     ScaffoldClusterOpts
		expectGolden string
	}{
		{
			name: "Should scaffold cluster declaration based on opts",
			withOpts: ScaffoldClusterOpts{
				Name:            "SomeClusterName",
				AWSAccountID:    "123456789012",
				Environment:     "production",
				Organization:    "oslokommune",
				RepositoryName:  "my_repo",
				OutputDirectory: "my_infrastructure",
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer

			err := ScaffoldClusterDeclaration(&buf, tc.withOpts)
			assert.NoError(t, err)

			g := goldie.New(t)
			g.Assert(t, tc.name, buf.Bytes())
		})
	}
}