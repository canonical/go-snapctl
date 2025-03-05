package snapctl_test

import (
	"testing"

	"github.com/canonical/go-snapctl"
	"github.com/stretchr/testify/require"
)

func TestInstall(t *testing.T) {
	t.Run("snapctl install --help valid-name", func(t *testing.T) {
		err := snapctl.Install("mock-component-name").Help().Run()
		require.NoError(t, err)
	})

	t.Run("snapctl install --help invalid-name", func(t *testing.T) {
		err := snapctl.Install("mock+component+name").Help().Run()
		require.Error(t, err)
	})
}
