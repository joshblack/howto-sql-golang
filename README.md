# `howto-sql-golang`

## Migrations

Creating a migration:

```bash
migrate create -ext sql -dir database/migrations create_foo_table
```

Running migrations:

```bash
migrate -path database/migrations -database postgres://postgres:postgres@localhost:5432/howto_dev?sslmode=disable up
```
