.PHONY: test install migrations migrate migrations hamsterwheel-migrations

hamsterwheel-migrations:
	go-bindata -pkg database -o cmd/hamsterwheel-server/internal/database/migrations.go -prefix "cmd/hamsterwheel-server/db/migrations/" cmd/hamsterwheel-server/db/migrations/...

migrations: hamsterwheel-migrations

install: migrations
	go install ./cmd/hamsterwheel-server

test: migrations install
	dropdb hamsterwheel-test; createdb hamsterwheel-test
	go test -v ./...
	forego run -e .env.test hamsterwheel-server