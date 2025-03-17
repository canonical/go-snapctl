package log

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
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

	loggingModeKey = "GO_SNAPCTL_LOGGING_MODE"
)

func init() {
	initialize()
}

// use a different function than init to allow testing
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

	exe, err := os.Executable()
	if err != nil {
		printErr("error getting the executable name:", err)
	}
	label := exe[strings.LastIndex(exe, "/")+1:]

	globalLogger, err = setupLogger(loggingMode, label, debug)
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

// SetLabel sets the label for log messages.
// The formatting depends on the used logging facility:
//
//	For standard streams, label is used as a "<label>: <message>"
//	For syslog, label is used as snap.<snap-instance-name>.<label> tag
//
// By default, label is set to the executable name.
//
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
