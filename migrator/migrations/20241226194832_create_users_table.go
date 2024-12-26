package migrations

import (
	"context"
	"database/sql"
	"log"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateUsersTable, downCreateUsersTable)
}

func upCreateUsersTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	log.Println("Running up migration for 20241226194832_create_users_table.go")

	query := `
		CREATE TABLE "public"."Users" (
    		"id" uuid NOT NULL DEFAULT gen_random_uuid(),
    		"created_at" timestamptz NOT NULL DEFAULT now(),
    		"updated_at" timestamptz NOT NULL DEFAULT now(),
    		"email" varchar(255) NOT NULL UNIQUE,
    		"password" varchar(255) NOT NULL,
    		PRIMARY KEY ("id")
		);
	`
	_, err := tx.ExecContext(ctx, query)
	return err
}

func downCreateUsersTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	log.Println("Running down migration for 20241226194832_create_users_table.go")

	query := `DROP TABLE "public"."Users"`
	_, err := tx.ExecContext(ctx, query)
	return err
}
