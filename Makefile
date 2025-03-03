
try:
	snapcraft try --use-lxd
	snap try prime

sync:
	cp -r $$(ls | egrep -v '^prime') prime/

test:
	sudo go-snapctl-tester.test \
		./ \
		./log \
		./snapctl \
		./env
