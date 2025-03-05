package snapctl_test

import (
	"testing"

	"github.com/canonical/go-snapctl"
	"github.com/stretchr/testify/require"
)

/*
TestInstall runs `snapctl install` with the `--help` argument to do a smoke test of this function.
It is only possible to do a real component installation for a snap that is in the snap store.
*/
func TestInstall(t *testing.T) {
	t.Run("snapctl install --help valid-name", func(t *testing.T) {
		err := snapctl.Install("mock-component-name").Help().Run()
		require.NoError(t, err)
	})

	t.Run("snapctl install --help one two three", func(t *testing.T) {
		err := snapctl.Install("one", "two", "three").Help().Run()
		require.NoError(t, err)
	})

	t.Run("snapctl install --help invalid-name", func(t *testing.T) {
		err := snapctl.Install("mock+component+name").Help().Run()
		require.Error(t, err)
	})
}
