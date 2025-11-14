# Database Migrations

This directory contains database migration files for managing schema changes.

## Migration File Naming Convention

Migration files should follow this naming pattern:

```
<timestamp>_<description>.up.sql      # For applying the migration
<timestamp>_<description>.down.sql    # For rolling back the migration
```

Example:
```
20251114120000_create_users_table.up.sql
20251114120000_create_users_table.down.sql
```

## Migration Tools

You can use tools like:
- [golang-migrate](https://github.com/golang-migrate/migrate)
- [goose](https://github.com/pressly/goose)
- [sql-migrate](https://github.com/rubenv/sql-migrate)

## Example: Using golang-migrate

### Install
```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### Create a new migration
```bash
migrate create -ext sql -dir migrations -seq create_users_table
```

### Run migrations
```bash
migrate -path migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" up
```

### Rollback migrations
```bash
migrate -path migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" down 1
```

## Migration Best Practices

1. **Always create both up and down migrations**
2. **Test migrations on a copy of production data**
3. **Keep migrations small and focused**
4. **Never modify existing migration files** - create new ones instead
5. **Use transactions** when possible
6. **Document breaking changes** in migration files

## Example Migration Files

See the example migration files in this directory for reference.
