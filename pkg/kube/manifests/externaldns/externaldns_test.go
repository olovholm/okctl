package externaldns_test

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	"github.com/oslokommune/okctl/pkg/config"

	"github.com/oslokommune/okctl/pkg/integration"
	"github.com/oslokommune/okctl/pkg/kube"
	"github.com/oslokommune/okctl/pkg/kube/manifests/externaldns"
	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/yaml"
)

func TestNew(t *testing.T) {
	testCases := []struct {
		name string
		ext  *externaldns.ExternalDNS
	}{
		{
			name: "Should work",
			ext:  externaldns.New("ABC123456", "test.oslo.systems"),
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t)

			dm := tc.ext.DeploymentManifest()
			db, err := yaml.Marshal(dm)
			assert.NoError(t, err)
			g.Assert(t, "deployment.yaml", db)

			crm := tc.ext.ClusterRoleManifest()
			crb, err := yaml.Marshal(crm)
			assert.NoError(t, err)
			g.Assert(t, "clusterrole.yaml", crb)

			crbm := tc.ext.ClusterRoleBindingManifest()
			crbb, err := yaml.Marshal(crbm)
			assert.NoError(t, err)
			g.Assert(t, "clusterrolebinding.yaml", crbb)
		})
	}
}

func TestExternalDNS(t *testing.T) {
	config.SkipUnlessIntegration(t)

	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment, as this test is flaky here.")
	}

	dir, err := ioutil.TempDir("", "externalDNS")
	assert.NoError(t, err)

	defer func() {
		_ = os.RemoveAll(dir)
	}()

	cluster := integration.NewKubernetesCluster()

	err = cluster.Create(2 * time.Minute)

	defer func() {
		err := cluster.Cleanup()
		assert.NoError(t, err)
	}()

	assert.NoError(t, err)

	kubeConfPath, err := cluster.KubeConfig()
	assert.NoError(t, err)

	testCases := []struct {
		name string
		ext  *externaldns.ExternalDNS
	}{
		{
			name: "Should work",
			ext:  externaldns.New("ABC1234567", "test.oslo.systems"),
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			k, err := kube.New(kube.NewFromKubeConfig(kubeConfPath))
			assert.NoError(t, err)

			// Need to add k.Watch to the output her, but that means bringing up localstack, I think?
			_, err = k.Apply(kube.Applier{Fn: tc.ext.CreateDeployment}, kube.Applier{Fn: tc.ext.CreateClusterRole}, kube.Applier{Fn: tc.ext.CreateClusterRoleBinding})
			assert.NoError(t, err)

			outputs, err := cluster.Debug("kube-system")
			assert.NoError(t, err)

			for title, data := range outputs {
				log.Println(title)
				log.Println(data)
			}
		})
	}
}
