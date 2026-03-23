package main

import (
	"github.com/PaulReznikov/pet-store/internal/config"
	"github.com/PaulReznikov/pet-store/pkg/logger"
	"github.com/rs/zerolog/log"
)

func main() {
	logger.Init()

	_, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	log.Info().Msg("Application started successfully")
}
