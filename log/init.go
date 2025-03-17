package log

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

var (
	debug       bool
	loggingMode string
)

const (
	// appMode is the default logging mode, suitable for logging from a snap app
	appMode = "app"
	// hookMode is suitable for logging from a snap hook
	hookMode = "hook"

	loggingModeKey = "GOSNAPCTL_LOG"
)

func init() {
	initialize()
}

// use a different function to allow testing
func initialize() {
	value, err := exec.Command("snapctl", "get", "debug").CombinedOutput()
	if err != nil {
		printErr(err)
		os.Exit(1)
	}
	debug = (string(bytes.TrimSpace(value)) == "true")

	loggingMode = os.Getenv(loggingModeKey)
	if loggingMode == "" {
		loggingMode = appMode
	}

	globalLogger, err = setupLogger(loggingMode, "", debug)
	if err != nil {
		printErr(err)
		os.Exit(1)
	}
}

func printErr(a ...any) {
	// By default, standard errors get collected as syslog with "snapd" tag.
	// Add prefix to distinguish these from other snapd logs.
	fmt.Fprintf(os.Stderr, "go-snapctl.log.init: %s\n", a...)
}

func setupLogger(mode, label string, debug bool) (l logger, err error) {
	if mode == appMode {
		l, err = newAppLogger(label, debug)
		if err != nil {
			return nil, fmt.Errorf("error creating app logger instance: %s", err)
		}
	} else if mode == hookMode {
		l, err = newHookLogger(label, debug)
		if err != nil {
			return nil, fmt.Errorf("error creating hook logger instance: %s", err)
		}
	} else {
		return nil, fmt.Errorf("unknown %s:", mode)
	}
	return
}

// SetLabel adds a label to log messages.
// The formatting depends on the used logging facility:
//
//	For standard streams, label is added as a prefix, separated by a colon and a space.
//	For syslog, label is added as a suffix to snap.<snap-instance-name>, separated by a dot.
//
// By default, no label is set.
// This function is NOT thread-safe. It should not be called concurrently with
// the other logging functions of this package.
func SetLabel(label string) {
	var err error
	// set up logger again with the new label
	globalLogger, err = setupLogger(loggingMode, label, debug)
	if err != nil {
		printErr(err)
		os.Exit(1)
	}
}
