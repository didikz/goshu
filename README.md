# GOSHU

Goshu (Go Shorten Url) is a simple application to shorten url and redirect to your original url

## DB Migration

Using [Golang Migrate](https://github.com/golang-migrate/migrate)

- Create Migrations `migrate create -ext sql -dir db/migrations -seq create_urls_table`
- Run migrations `migrate -database "$(cat config.yml | python -c "import yaml,sys; print(yaml.safe_load(sys.stdin)['database']['url'])")" -source file://db/migrations up`

## Running The Service

- `go run ./cmd/goshu/main.go`
