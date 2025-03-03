build:
	snapcraft

install:
	sudo snap install --dangerous ./go-snapctl-tester_test_amd64.snap

# This was meant to be used to sync the prime directory with the source code in snapcraft try mode.
# It can't be used that way because snapcraft try doesn't work with core24.
# To make this method work, the test files should be kepts in a writable directory and copied over.
# sync:
# 	cp -r $$(ls | egrep -v '^prime') prime/

test:
	sudo go-snapctl-tester.test \
		./ \
		./log \
		./env
