build:
	snapcraft -v

clean-build:
	snapcraft clean
	make build

install:
	sudo snap install --dangerous ./go-snapctl-tester_test_amd64.snap
	sudo snap connect go-snapctl-tester:home

test:
	sudo go-snapctl-tester.test ./...

remove:
	sudo snap remove go-snapctl-tester
