run:
  go run .

test:
  go test ./...

up:
  goose -dir "sql/schema" postgres "postgres://pd:@localhost:5432/gator" up

down:
  goose -dir "sql/schema" postgres "postgres://pd:@localhost:5432/gator" down
