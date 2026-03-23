package main

import (
	"context"
	"embed"
	"log"
	"os"

	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

//go:embed scripts/*.sql
var embedMigrations embed.FS

func main() {
	goose.SetBaseFS(embedMigrations)

	dbString := os.Getenv("DATABASE_URL")
	if dbString == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := sql.Open("pgx", dbString)
	if err != nil {
		log.Fatalf("Failed to open DB: %v\n", err)
	}
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Failed to set dialect: %v\n", err)
	}

	command := "up"
	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	log.Printf("Running goose command: '%s'...\n", command)

	if err := goose.RunContext(context.Background(), command, db, "scripts"); err != nil {
		log.Fatalf("goose %s failed: %v\n", command, err)
	}

	log.Println("Migration command completed successfully!")
}
