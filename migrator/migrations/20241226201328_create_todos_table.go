package migrations

import (
	"context"
	"database/sql"
	"log"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTodosTable, downCreateTodosTable)
}

func upCreateTodosTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	log.Println("Running up migration for 20241226201328_create_todos_table.go")

	todosTableCreateQuery := `
		CREATE TABLE "public"."Todos" (
    		"id" uuid NOT NULL DEFAULT gen_random_uuid(),
    		"created_at" timestamptz NOT NULL DEFAULT now(),
    		"updated_at" timestamptz NOT NULL DEFAULT now(),
    		"user_id" uuid NOT NULL,
    		"title" varchar(255) NOT NULL,
			"status" "enum_Todos_status" NOT NULL DEFAULT 'PENDING',
    		PRIMARY KEY ("id"),
    		FOREIGN KEY ("user_id") REFERENCES "public"."Users" ("id") ON DELETE CASCADE
		);
	`
	_, err := tx.ExecContext(ctx, todosTableCreateQuery)
	return err
}

func downCreateTodosTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	log.Println("Running down migration for 20241226201328_create_todos_table.go")

	todosTableDropQuery := `DROP TABLE "public"."Todos"`
	_, err := tx.ExecContext(ctx, todosTableDropQuery)
	return err
}
