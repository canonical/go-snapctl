package snapctl_test

import (
	"reflect"
	"testing"

	"github.com/canonical/go-snapctl"
	"github.com/stretchr/testify/require"
)

func TestServices(t *testing.T) {

	t.Run("snapctl services", func(t *testing.T) {
		t.Run("one", func(t *testing.T) {
			services, err := snapctl.Services(mockService).Run()
			require.NoError(t, err, "Error getting services.")
			require.Len(t, services, 1)
			serviceName := reflect.ValueOf(services).MapKeys()[0].String()
			require.Equal(t, mockService, serviceName)
		})

		t.Run("all", func(t *testing.T) {
			services, err := snapctl.Services().Run()
			require.NoError(t, err, "Error getting services.")
			require.Len(t, services, 2)
			for k := range services {
				require.Contains(t, []string{mockService, mockService2}, k)
			}
		})

		t.Run("enabled and active", func(t *testing.T) {
			startAndEnableService(t, mockService)
			t.Cleanup(func() { stopAndDisableService(t, mockService) })

			services, err := snapctl.Services(mockService).Run()
			require.NoError(t, err, "Error getting services.")
			require.Len(t, services, 1)
			v, found := services[mockService]
			require.True(t, found)
			// Startup and Current will only be these values on an English host machine
			require.Equal(t, v.Startup, "enabled", "Service not enabled")
			require.Equal(t, v.Current, "active", "Service not active")
		})

		t.Run("disabled and inactive", func(t *testing.T) {
			services, err := snapctl.Services(mockService2).Run()
			require.NoError(t, err, "Error getting services.")
			require.Len(t, services, 1)
			v, found := services[mockService2]
			require.True(t, found)
			// Startup and Current will only be these values on an English host machine
			require.Equal(t, v.Startup, "disabled", "Service not disabled")
			require.Equal(t, v.Current, "inactive", "Service not inactive")
		})

		t.Run("service not found", func(t *testing.T) {
			_, err := snapctl.Services("non-existed").Run()
			require.Error(t, err)
		})

		t.Run("reject service name with space", func(t *testing.T) {
			_, err := snapctl.Services("bad name").Run()
			require.Error(t, err)
		})

	})
}
