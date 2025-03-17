# go-snapctl log

This package provides utilities for snap app and hook logging.

The logging mode can be changed via the `GO_SNAPCTL_LOGGING_MODE` environment variable.
Supported logging modes are:
- `app` (default) — using standard streams
- `hook` — using syslog with the addition of standard output for errors

Debug logging can be enable by setting the `debug=true` snap option.

For more usage instructions, refer to the package reference: https://pkg.go.dev/github.com/canonical/go-snapctl/log 
