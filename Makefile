install-tools:
	go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/golang/mock/mockgen@v1.6.0

migrate-create:
	migrate create -ext sql -dir src/database/migrations -seq ${name}

migrate:
	go run src/server/main.go migrate-up

migrate-down:
	go run src/server/main.go migrate-down