# Database Migrations

This directory contains SQLite database migration files for managing schema changes.

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
- [golang-migrate](https://github.com/golang-migrate/migrate) - Recommended
- [goose](https://github.com/pressly/goose)
- [sql-migrate](https://github.com/rubenv/sql-migrate)

## Example: Using golang-migrate with SQLite

### Install
```bash
go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### Create a new migration
```bash
migrate create -ext sql -dir migrations -seq add_posts_table
```

### Run migrations
```bash
migrate -path migrations -database "sqlite3://data/app.db" up
```

### Rollback migrations
```bash
migrate -path migrations -database "sqlite3://data/app.db" down 1
```

### Check migration version
```bash
migrate -path migrations -database "sqlite3://data/app.db" version
```

## SQLite-Specific Notes

### Data Types
SQLite has a flexible type system. Common mappings:
- `INTEGER` - Integer numbers (can be PRIMARY KEY AUTOINCREMENT)
- `TEXT` - String data (replaces VARCHAR, CHAR)
- `REAL` - Floating point numbers
- `BLOB` - Binary data
- `DATETIME` - Date and time (stored as TEXT, INTEGER, or REAL)

### Constraints
SQLite supports:
- PRIMARY KEY
- UNIQUE
- NOT NULL
- CHECK
- DEFAULT
- FOREIGN KEY (must be enabled with PRAGMA)

### Indexes
Create indexes using:
```sql
CREATE INDEX IF NOT EXISTS idx_name ON table(column);
```

## Migration Best Practices

1. **Always create both up and down migrations**
2. **Test migrations on a copy of production data**
3. **Keep migrations small and focused**
4. **Never modify existing migration files** - create new ones instead
5. **Use transactions** when possible (SQLite does this automatically for most operations)
6. **Document breaking changes** in migration files
7. **Use IF NOT EXISTS** for idempotent migrations

## Example Migration Files

See the example migration files in this directory for reference. The included example creates a users table with proper indexes.

## SQLite Database Location

By default, the database is stored at `data/app.db`. You can change this via the `DB_PATH` environment variable.

## WAL Mode (Recommended)

For better concurrency, enable WAL (Write-Ahead Logging) mode:

```sql
PRAGMA journal_mode=WAL;
```

This is recommended for production use and allows concurrent readers during writes.
