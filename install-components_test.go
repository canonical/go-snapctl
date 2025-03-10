package snapctl_test

import (
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
		err := snapctl.InstallComponents("valid-name").Run()
		require.ErrorContainsf(t, err, "cannot install components for a snap that is unknown to the store", "Unexpected error returned")
	})

	t.Run("snapctl install +one +two +three", func(t *testing.T) {
		err := snapctl.InstallComponents("one", "two", "three").Run()
		require.ErrorContainsf(t, err, "cannot install components for a snap that is unknown to the store", "Unexpected error returned")
	})

	t.Run("snapctl install +invalid name", func(t *testing.T) {
		err := snapctl.InstallComponents("invalid name").Run()
		require.ErrorContainsf(t, err, "component names must not contain spaces", "Unexpected error returned")
	})

	t.Run("snapctl install +<no name>", func(t *testing.T) {
		err := snapctl.InstallComponents().Run()
		require.ErrorContainsf(t, err, "at least one component must be specified", "Unexpected error returned")
	})
}
