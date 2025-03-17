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
	// hookMode is the default logging mode, suitable for logging from a snap hook
	hookMode = "hook"
	// appMode is suitable for logging from a snap app
	appMode = "app"
)

func init() {
	initialize()
}

// use a different function to allow testing
func initialize() {
	value, err := exec.Command("snapctl", "get", "debug").CombinedOutput()
	if err != nil {
		stderr(err)
		os.Exit(1)
	}
	debug = (string(bytes.TrimSpace(value)) == "true")

	loggingMode = os.Getenv("LOGGING_MODE")
	if loggingMode == "" || loggingMode == hookMode {
		globalLogger, err = newHookLogger("", debug)
		if err != nil {
			stderr("error creating syslog instance: %s", err)
			os.Exit(1)
		}
	} else if loggingMode == appMode {
		globalLogger, err = newAppLogger("", debug)
		if err != nil {
			stderr("error creating app logger instance: %s", err)
			os.Exit(1)
		}
	} else {
		stderr("unknown LOGGING_MODE: %s", loggingMode)
		os.Exit(1)
	}
}

func stderr(a ...any) {
	// Standard errors get collected with "snapd" as syslog app.
	// We add the tag as prefix to distinguish these from other snapd logs.
	fmt.Fprintf(os.Stderr, "go-snapctl.log.init: %s\n", a...)
}

// SetLabel adds a label to log messages.
// The formatting depends on the logging mode.
// By default, no label is set.
// This function is NOT thread-safe. It should not be called concurrently with
// the other logging functions of this package.
func SetLabel(label string) {
	var err error
	if loggingMode == hookMode {
		globalLogger, err = newHookLogger(label, debug)
		if err != nil {
			stderr("error creating syslog instance: %s", err)
			os.Exit(1)
		}
	} else if loggingMode == appMode {
		globalLogger, err = newAppLogger(label, debug)
		if err != nil {
			stderr("error creating app logger instance: %s", err)
			os.Exit(1)
		}
	} else {
		stderr("unknown LOGGING_MODE: %s", loggingMode)
		os.Exit(1)
	}
}
