package snapctl_test

import (
	"testing"

	"github.com/canonical/go-snapctl"
	"github.com/stretchr/testify/require"
)

/*
TestRemove smoke tests the component removal.
It is only possible to do a real component removal for a snap that is in the snap store.
*/
func TestRemove(t *testing.T) {
	t.Run("snapctl remove +valid-name", func(t *testing.T) {
		err := snapctl.RemoveComponents("valid-name").Run()
		require.ErrorContainsf(t, err, "component \"valid-name\" is not installed for revision", "Unexpected error returned")
	})

	t.Run("snapctl remove +one +two +three", func(t *testing.T) {
		err := snapctl.RemoveComponents("one", "two", "three").Run()
		require.ErrorContainsf(t, err, "component \"one\" is not installed for revision", "Unexpected error returned")
	})

	t.Run("snapctl remove +invalid name", func(t *testing.T) {
		err := snapctl.RemoveComponents("invalid name").Run()
		require.ErrorContainsf(t, err, "component names must not contain spaces", "Unexpected error returned")
	})

	t.Run("snapctl remove +<no name>", func(t *testing.T) {
		err := snapctl.RemoveComponents().Run()
		require.ErrorContainsf(t, err, "at least one component must be specified", "Unexpected error returned")
	})
}
