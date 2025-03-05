/*
Usage:
  snapctl [OPTIONS] install <snap|snap+comp|+comp>...

The install command installs components.


Help Options:
  -h, --help                        Show this help message

[install command arguments]
  <snap|snap+comp|+comp>:           Components to be installed (snap must be
                                    the caller snap if specified).
*/

package snapctl

import (
	"fmt"
	"strings"
)

type install struct {
	components []string
	options    []string
	validators []func() error
}

// Install installs components of the snap
// It takes an arbitrary number of component names as input, and should be in the format `snap+component` or `+component`
// It returns an object for setting the CLI arguments before running the command
func Install(components ...string) (cmd install) {
	cmd.components = components

	cmd.validators = append(cmd.validators, func() error {
		for _, key := range cmd.components {
			if strings.Contains(key, " ") {
				return fmt.Errorf("component names must not contain spaces. Got: '%s'", key)
			}
		}
		return nil
	})

	return cmd
}

// Run executes the start command
func (cmd install) Run() error {
	// validate all input
	for _, validate := range cmd.validators {
		if err := validate(); err != nil {
			return err
		}
	}

	// construct the command args
	// install [install-OPTIONS] <snap|snap+comp|+comp>...
	var args []string
	// options
	args = append(args, cmd.options...)
	// components
	args = append(args, cmd.components...)

	_, err := run("install", args...)
	return err
}
