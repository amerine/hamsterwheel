.PHONY: test

install:
	go install ./cmd/hamsterwheel-server

test: install
	dropdb hamsterwheel-test; createdb hamsterwheel-test
	go test -v ./...
	forego run -e .env.test hamsterwheel-server