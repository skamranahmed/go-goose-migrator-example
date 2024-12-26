package migrations

import (
	"context"
	"database/sql"
	"log"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTodosStatusEnum, downCreateTodosStatusEnum)
}

func upCreateTodosStatusEnum(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	log.Println("Running up migration for 20241226201153_create_todos_status_enum.go")

	statusEnumCreateQuery := `
		CREATE TYPE "enum_Todos_status" AS ENUM ('PENDING', 'IN_PROGRESS', 'DONE');
	`
	_, err := tx.ExecContext(ctx, statusEnumCreateQuery)
	return err
}

func downCreateTodosStatusEnum(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	log.Println("Running down migration for 20241226201153_create_todos_status_enum.go")

	statusEnumDropQuery := `DROP TYPE "enum_Todos_status"`
	_, err := tx.ExecContext(ctx, statusEnumDropQuery)
	return err
}
