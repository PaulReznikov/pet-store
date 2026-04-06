package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/PaulReznikov/pet-store/internal/config"
)

const formatJSON = "json"

func New(logConf *config.Logger) zerolog.Logger {
	level := zerolog.InfoLevel

	// err == nil
	if newLevel, err := zerolog.ParseLevel(logConf.LogLevel); err == nil {
		level = newLevel
	}

	var out io.Writer = os.Stdout

	if logConf.LogFormat != formatJSON {
		out = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
			NoColor:    logConf.LogNoColor,
		}
	}

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	return zerolog.New(out).Level(level).With().Timestamp().Caller().Logger()
}
