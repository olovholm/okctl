package storm_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/oslokommune/okctl/pkg/breeze"
	"github.com/oslokommune/okctl/pkg/client"

	"github.com/oslokommune/okctl/pkg/client/mock"

	"github.com/oslokommune/okctl/pkg/client/core/state/storm"
	"github.com/stretchr/testify/assert"
)

func TestDomainStateScenario(t *testing.T) {
	dir, err := ioutil.TempDir(os.TempDir(), "storm")
	assert.NoError(t, err)

	defer func() {
		err = os.RemoveAll(dir)
		assert.NoError(t, err)
	}()

	db := breeze.New(filepath.Join(dir, "storm.db"))

	err = db.Init(&storm.HostedZone{})
	assert.NoError(t, err)

	state := storm.NewDomainState(db)

	err = state.SaveHostedZone(mock.HostedZone())
	assert.NoError(t, err)

	hz, err := state.GetHostedZone(mock.DefaultDomain)
	assert.NoError(t, err)
	assert.Equal(t, mock.HostedZone(), hz)

	hz.NameServers = nil
	err = state.UpdateHostedZone(hz)
	assert.NoError(t, err)

	hzs, err := state.GetHostedZones()
	assert.NoError(t, err)
	assert.Equal(t, []*client.HostedZone{hz}, hzs)

	hz.FQDN = "fake5"
	err = state.SaveHostedZone(hz)
	assert.NoError(t, err)

	err = state.RemoveHostedZone(mock.DefaultDomain)
	assert.NoError(t, err)

	err = state.RemoveHostedZone(mock.DefaultDomain)
	assert.NoError(t, err)
}
