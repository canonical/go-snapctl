// Utility testing functions

package snapctl_test

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	snapName     = "go-snapctl-tester"
	mockService  = snapName + ".mock-service"
	mockService2 = snapName + ".mock-service-2"
)

func setConfigValue(t *testing.T, key, value string) {
	output, err := exec.Command("snapctl", "set", fmt.Sprintf("%s=%s", key, value)).CombinedOutput()
	assert.NoError(t, err,
		"Error setting config value via snapctl: %s", output)
}

func getConfigStrictValue(t *testing.T, key string) string {
	output, err := exec.Command("snapctl", "get", "-t", key).CombinedOutput()
	require.NoError(t, err,
		"Error getting config value via snapctl: %s", output)
	return strings.TrimSpace(string(output))
}

/*
getServiceStatus uses snapctl to obtain the Startup and Current states of the given service.
The response is assumed to be English, which can not be guaranteed in all environments.
But this function is only used in tests where the environment is controlled.
*/
func getServiceStatus(t *testing.T, service string) (enabled, active bool) {
	output, err := exec.Command("snapctl", "services", service).CombinedOutput()
	require.NoError(t, err,
		"Error getting services via snapctl: %s", output)

	// Startup and Current will only be these values on an English host machine
	enabled = strings.Contains(string(output), "enabled")
	// look for not "inactive", because both active and inactive contain "active"
	active = !strings.Contains(string(output), "inactive")
	return enabled, active
}

func startService(t *testing.T, service string) {
	output, err := exec.Command("snapctl", "start", service).CombinedOutput()
	require.NoError(t, err,
		"Error starting service via snapctl: %s", output)
}

func startAndEnableService(t *testing.T, service string) {
	output, err := exec.Command("snapctl", "start", "--enable", service).CombinedOutput()
	require.NoError(t, err,
		"Error starting service via snapctl: %s", output)
}

func stopAndEnableAllServices(t *testing.T) {
	startAndEnableService(t, snapName)
}

func stopAndDisableService(t *testing.T, service string) {
	output, err := exec.Command("snapctl", "stop", "--disable", service).CombinedOutput()
	require.NoError(t, err,
		"Error stopping service via snapctl: %s", output)
}

func stopAndDisableAllServices(t *testing.T) {
	stopAndDisableService(t, snapName)
}
