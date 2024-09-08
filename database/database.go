package database

import (
	"database/sql"
	"embed"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

//go:embed sql_migration/*.sql
var dbMigrations embed.FS

var DbConnection *sql.DB

func DBMigrate(dbParam *sql.DB, direction string) {
	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: dbMigrations,
		Root:       "sql_migration",
	}

	var migrateDirection migrate.MigrationDirection
	switch direction {
	case "up":
		migrateDirection = migrate.Up
	case "down":
		migrateDirection = migrate.Down
	default:
		fmt.Println("Invalid migration direction. Use 'up' or 'down'.")
		return
	}

	n, errs := migrate.Exec(dbParam, "postgres", migrations, migrateDirection)
	if errs != nil {
		panic(errs)
	}

	DbConnection = dbParam

	fmt.Println("Migration success, applied", n, "migrations!")
}
