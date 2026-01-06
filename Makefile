.DEFAULT_GOAL := all

all: build install test

build:
	snapcraft pack -v

clean-build:
	snapcraft clean
	make build

install:
	sudo snap install --dangerous ./go-snapctl-tester_test_amd64.snap
	sudo snap connect go-snapctl-tester:home
	sudo snap connect go-snapctl-tester:system-observe

test:
	sudo go-snapctl-tester.test ./...

remove:
	sudo snap remove go-snapctl-tester

clean: remove
	snapcraft clean
