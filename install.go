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
	snap       string
	components []string
	options    []string
	validators []func(install) error
}

// Install creates and returns an instance of the install command.
func Install() (cmd install) {
	cmd.validators = append(cmd.validators, func(cmd install) error {
		if len(cmd.components) == 0 {
			return fmt.Errorf("at least one component must be specified")
		}
		return nil
	})

	cmd.validators = append(cmd.validators, func(cmd install) error {
		for _, component := range cmd.components {
			if strings.Contains(component, " ") {
				return fmt.Errorf("component names must not contain spaces. Got: '%s'", component)
			}
		}
		return nil
	})

	return cmd
}

// Snap adds the snap name that should be installed, or for which the components will be installed
func (cmd install) Snap(snap string) install {
	cmd.snap = snap
	return cmd
}

// Components adds the list of components that will be installed for the current snap, or the one set by Snap()
func (cmd install) Components(components ...string) install {
	cmd.components = append(cmd.components, components...)
	return cmd
}

// Run executes the install command
func (cmd install) Run() error {
	// validate all input
	for _, validate := range cmd.validators {
		if err := validate(cmd); err != nil {
			return err
		}
	}

	// construct the command args
	// install [install-OPTIONS] <snap|snap+comp|+comp>...
	var args []string
	// options
	args = append(args, cmd.options...)

	// If a snap name is set, use it as a prefix for the components. Always prefix component names with a +.
	for _, component := range cmd.components {
		args = append(args, cmd.snap+"+"+component)
	}

	_, err := run("install", args...)
	return err
}
