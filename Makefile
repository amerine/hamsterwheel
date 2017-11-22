.PHONY: test

test:
	dropdb hamsterwheel-test; createdb hamsterwheel-test
	go test -v ./...