package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"

	_ "github.com/skamranahmed/go-goose-migrator-example/migrator/migrations"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		log.Fatal("Please provide a Goose command")
	}

	dbUser := os.Getenv("DATABASE_USER")
	if dbUser == "" {
		log.Fatal("Database user name cannot be empty")
	}

	dbPassword := os.Getenv("DATABASE_PASSWORD")
	if dbPassword == "" {
		log.Fatal("Database password cannot be empty")
	}

	dbName := os.Getenv("DATABASE_NAME")
	if dbName == "" {
		log.Fatal("Database name cannot be empty")
	}

	dbHost := os.Getenv("DATABASE_HOST")
	if dbHost == "" {
		log.Fatal("Database host cannot be empty")
	}

	dbPort, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if dbPort == 0 {
		log.Fatal("Database port cannot be empty")
	}

	if err != nil {
		log.Fatalf("Invalid DATABASE_PORT: %+v", err)
	}

	dbDsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable", dbUser, dbPassword, dbName, dbHost, dbPort)

	db, err := sql.Open("postgres", dbDsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %+v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("unable to establish connection with the db, error: %+v", err)
	}

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Failed to set dialect: %+v", err)
	}

	goose.SetTableName(`"GooseMigrationsMetadata"`)

	command := args[0]
	migrationDir := "migrations"

	if command == "up" {
		err = goose.UpContext(context.Background(), db, migrationDir, goose.WithAllowMissing())
		if err != nil {
			log.Fatalf("Failed to run Goose command: %+v", err)
		}
	} else {
		err = goose.RunContext(context.Background(), command, db, migrationDir, args[1:]...)
		if err != nil {
			log.Fatalf("Failed to run Goose command: %+v", err)
		}
	}
}
