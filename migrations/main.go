package main

import (
	"database/sql"
	"embed"
	"os"

	"github.com/PaulReznikov/pet-store/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//go:embed scripts/*.sql
var embedMigrations embed.FS

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	goose.SetBaseFS(embedMigrations)
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}
	dbString := config.Postgres.CreateDBConnectionString()
	if dbString == "" {
		log.Fatal().Msg("DATABASE_URL is not set")
	}

	db, err := sql.Open("pgx", dbString)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open DB")
	}
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal().Err(err).Msg("Failed to set dialect")
	}

	currentVersion, err := goose.GetDBVersion(db)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get current DB version")
	}

	log.Info().Int64("current_version", currentVersion).Msg("Starting migration")

	if err := goose.Up(db, "scripts"); err != nil {
		log.Error().Err(err).Msg("Migration failed, attempting rollback")

		if rollbackErr := goose.DownTo(db, "scripts", currentVersion); rollbackErr != nil {
			log.Fatal().
				Err(err).
				AnErr("rollback_error", rollbackErr).
				Msg("Migration failed AND rollback failed")
		}

		log.Fatal().Err(err).Msg("Migration failed (rollback succeeded)")
	}

	newVersion, _ := goose.GetDBVersion(db)
	log.Info().
		Int64("from_version", currentVersion).
		Int64("to_version", newVersion).
		Msg("Migration completed successfully")
}
