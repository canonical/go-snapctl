package snapctl

import (
	"context"
	"fmt"
	"os"

	"github.com/coreos/go-systemd/v22/dbus"
	godbus "github.com/godbus/dbus/v5" // Import the underlying dbus library
)

// IsRunning returns the state of a given service.
// serviceName should be in the format snap.<snap-name>.<app-name>.service
func IsRunning(serviceName string) (bool, error) {
	// Ensure the service name is in the correct format for systemd
	serviceName = "snap." + serviceName + ".service"

	// 1. Explicitly connect to the System Bus (allowed by system-observe)
	godbusConn, err := godbus.SystemBus()
	if err != nil {
		return false, fmt.Errorf("error connecting to DBus system bus: %v", err)
	}

	// 2. Wrap the connection with go-systemd to get the convenient helper methods
	conn, err := dbus.NewConnection(func() (*godbus.Conn, error) {
		return godbusConn, nil
	})
	if err != nil {
		return false, fmt.Errorf("error initializing systemd client: %v", err)
	}
	defer conn.Close()

	// 3. Get Unit Properties
	props, err := conn.GetUnitPropertiesContext(context.Background(), serviceName)
	if err != nil {
		fmt.Println("Error fetching status:", err)
		os.Exit(1)
	}

	/*
		    fmt.Printf("Service: %s\n", serviceName)
			fmt.Printf("ActiveState: %s\n", props["ActiveState"]) // "active", "inactive", "failed"
			fmt.Printf("LoadState: %s\n", props["LoadState"])     // "loaded", "not-found"
	*/

	if activeState, ok := props["ActiveState"].(string); ok {
		switch activeState {
		case "active":
			return true, nil
		case "inactive":
			return false, nil
		default:
			return false, fmt.Errorf("service %s is in an unexpected state: %s", serviceName, activeState)
		}
	}

	return false, fmt.Errorf("could not determine the state of service %s", serviceName)
}
