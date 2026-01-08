// Utility testing functions

package snapctl_test

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/canonical/go-snapctl"
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
	services, err := snapctl.Services(service).Run()
	require.NoError(t, err, "Error getting services: %v", err)

	serviceStatus, found := services[service]
	require.True(t, found, "Service %s not found", service)

	// validate the Startup value
	if serviceStatus.Startup != "enabled" && serviceStatus.Startup != "disabled" {
		t.Fatalf("unexpected snapctl output: expected Startup as enabled|disabled, got: %s", serviceStatus.Startup)
	}

	// validate the Current value
	if serviceStatus.Current != "active" && serviceStatus.Current != "inactive" {
		t.Fatalf("unexpected snapctl output: expected Current as active|inactive, got: %s", serviceStatus.Current)
	}

	// Startup and Current will only be these values on an English host machine
	enabled = serviceStatus.Startup == "enabled"
	active = serviceStatus.Current == "active"
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

// isEnglishLocale checks if the system's locale is likely English
// TODO remove this function and related checks when snapctl supports locale-independent output
// See bug https://bugs.launchpad.net/snapd/+bug/2137543
func isEnglishLocale() bool {
	// Common environment variables for locale on Unix/Linux/macOS
	lang := os.Getenv("LANG")
	lcAll := os.Getenv("LC_ALL")

	// Check if any of these indicate an English locale (e.g., "en_US", "en-GB", "en")
	if strings.HasPrefix(strings.ToLower(lang), "en") || strings.HasPrefix(strings.ToLower(lcAll), "en") {
		return true
	}

	return false
}
