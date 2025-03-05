package snapctl_test

import (
	"os"
	"testing"

	"github.com/canonical/go-snapctl"
	"github.com/stretchr/testify/require"
)

/*
TestInstall runs `snapctl install` to do a smoke test of this function.
It is only possible to do a real component installation for a snap that is in the snap store.
*/
func TestInstall(t *testing.T) {
	t.Run("snapctl install +valid-name", func(t *testing.T) {
		err := snapctl.Install("+valid-name").Run()
		require.ErrorContainsf(t, err, "cannot install components for a snap that is unknown to the store", "returned expected error")
	})

	t.Run("snapctl install +one +two +three", func(t *testing.T) {
		err := snapctl.Install("+one", "+two", "+three").Run()
		require.ErrorContainsf(t, err, "cannot install components for a snap that is unknown to the store", "returned expected error")
	})

	t.Run("snapctl install +invalid name", func(t *testing.T) {
		err := snapctl.Install("+invalid name").Run()
		require.ErrorContainsf(t, err, "component names must not contain spaces", "returned expected error")
	})

	t.Run("snapctl install $SNAP_NAME", func(t *testing.T) {
		err := snapctl.Install(os.Getenv("SNAP_NAME")).Run()
		require.NoError(t, err)
	})

	t.Run("snapctl install invalid-snap", func(t *testing.T) {
		err := snapctl.Install("invalid-snap").Run()
		require.ErrorContainsf(t, err, "cannot install snaps using snapctl", "returned expected error")
	})
}
