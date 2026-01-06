package snapctl_test

import (
	"testing"

	"github.com/canonical/go-snapctl"
	"github.com/stretchr/testify/require"
)

func TestIsRunning(t *testing.T) {
	t.Cleanup(func() { stopAndDisableAllServices(t) })

	err := snapctl.Start(mockService).Run()
	require.NoError(t, err)

	active, err := snapctl.IsRunning(mockService)
	require.NoError(t, err)
	require.True(t, active, "active")

	active, err = snapctl.IsRunning(mockService2)
	require.NoError(t, err)
	require.False(t, active, "inactive")
}
