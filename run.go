package snapctl

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/canonical/go-snapctl/log"
)

func run(subcommand string, subargs ...string) (string, error) {
	args := []string{subcommand}
	args = append(args, subargs...)

	log.Debugf("Executing 'snapctl %s'\n", strings.Join(args, " "))

	cmd := exec.Command("snapctl", args...)
	cmd.Env = os.Environ()

	// Set locale to C to ensure output is in English
	// This has no effect, as the host locale is being used by snapd, and this env var does not propagate across the REST API
	cmd.Env = append(cmd.Env, "LANG=C.utf8")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%w: %s", err, output)
	}

	return strings.TrimSpace(string(output)), nil
}
