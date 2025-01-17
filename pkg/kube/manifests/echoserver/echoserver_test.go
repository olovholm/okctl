package echoserver_test

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	"github.com/oslokommune/okctl/pkg/config"

	"github.com/oslokommune/okctl/pkg/integration"
	"github.com/oslokommune/okctl/pkg/kube"
	"github.com/oslokommune/okctl/pkg/kube/manifests/echoserver"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

//nolint:funlen
func TestEchoServer(t *testing.T) {
	config.SkipUnlessIntegration(t)

	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment, as this test is flaky here.")
	}

	dir, err := ioutil.TempDir("", "echoserver")
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
		ext  *echoserver.EchoServer
	}{
		{
			name: "Should work",
			ext:  echoserver.New("default"),
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			k, err := kube.New(kube.NewFromKubeConfig(kubeConfPath))
			assert.NoError(t, err)

			k.WithLogger(logrus.StandardLogger())

			resources, err := k.Apply(kube.Applier{Fn: tc.ext.CreateDeployment})
			assert.NoError(t, err)

			err = k.Watch(resources, 2*time.Minute)
			assert.NoError(t, err)

			outputs, err := cluster.Debug("default")
			assert.NoError(t, err)

			for title, data := range outputs {
				log.Println(title)
				log.Println(data)
			}
		})
	}
}
