# go-snapctl
[![Go Reference](https://pkg.go.dev/badge/github.com/canonical/go-snapctl.svg)](https://pkg.go.dev/github.com/canonical/go-snapctl)

Go wrapper library for the [snapctl](https://snapcraft.io/docs/using-snapctl) tool.

Wrappers for following subcommands are partially implemented for EdgeX use cases:

- [ ] `fde-setup-request`: Obtain full disk encryption setup request
- [ ] `fde-setup-result`: Set result for full disk encryption
- [x] `get`: The get command prints configuration and interface connection settings.                
- [x] `is-connected`: Return success if the given plug or slot is connected, and failure otherwise
- [ ] `reboot`: Control the reboot behavior of the system          
- [x] `restart`: Restart services    
- [x] `services`: Query the status of services      
- [x] `set`: Changes configuration options
- [ ] `set-health`: Report the health status of a snap
- [x] `start`: Start services 
- [x] `stop`: Stop services
- [ ] `system-mode`: Get the current system mode and associated details
- [x] `unset`: Remove configuration options

The commands and descriptions are from `snapctl --help`.

# Usage

```go
package main

import (
	"fmt"

	"github.com/canonical/go-snapctl/snapctl"
)

func main() {
	// unset
	err := snapctl.Unset("http").Run()
	if err != nil {
		panic(err)
	}

	// set values
	err = snapctl.Set("http.bind-address", "0.0.0.0").Run()
	if err != nil {
		panic(err)
	}
	err = snapctl.Set("http.bind-port", "8080").Run()
	if err != nil {
		panic(err)
	}

	// set values with a JSON object
	err = snapctl.Set("http.tls",
		`{
			"enabled":"true",
			"cert":"./cert.pem",
			"privkey":"./key.pem"
		}`).Document().Run()
	if err != nil {
		panic(err)
	}

	// get one value
	value, err := snapctl.Get("http.bind-port").Run()
	if err != nil {
		panic(err)
	}
	fmt.Println(value)
	// Outputs:
	// 8080

	// get values as JSON object
	value, err = snapctl.Get("http").Document().Run()
	if err != nil {
		panic(err)
	}
	fmt.Println(value)
	// Outputs:
	// {
	//   "bind-address": "0.0.0.0",
	//   "bind-port": 8080,
	//   "tls": {
	//     "cert": "./cert.pem",
	// 	   "enabled": "true",
	// 	   "privkey": "./key.pem"
	//   }
	// }

	// start and enable a service
	err := snapctl.Start("snap-name.service-name").Enable().Run()
	if err != nil {
		panic(err)
	}

	// start all services
	err := snapctl.Start("snap.name").Run()
	if err != nil {
		panic(err)
	}
}
```

### Testing
The tests need to run in a snap environment:

Build and install:
```bash
snapcraft
sudo snap install --dangerous ./go-snapctl-tester_test_amd64.snap
```

The tests files are read relative to project source inside the snap.
The `go-snapctl-tester.test` command runs `go test -v --cover` internally and accepts
all other go test arguments.

Run all tests:
```
make test
```

Run top-level tests:
```bash
sudo go-snapctl-tester.test
```

Run tests in one package, e.g. `snapctl`:
```bash
sudo go-snapctl-tester.test ./snapctl
```

Run one unit test, e.g. `TestGet`:
```bash
sudo go-snapctl-tester.test ./snapctl -run TestGet
```

#### Development
```
make try
```

You can now edit the files locally, copy them to prime directory, and re-run the
tests without rebuilding the project. E.g.:

```
make sync
sudo go-snapctl-tester.test ./snapctl
```
