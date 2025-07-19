//go:build !nosnap

package env

import (
	"os"

	"github.com/canonical/go-snapctl/log"
)

func init() {
	if err := getEnvVars(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
