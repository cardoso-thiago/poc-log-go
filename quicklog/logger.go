package quicklog

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("applicationProperties")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Error().Err(err).Msg("Erro ao ler o arquivo de configuração")
	}

	zerolog.TimeFieldFormat = "02/01/2006 15:04:05.000"
	//Apenas para validar a característica de mudança dos nomes dos campos
	zerolog.TimestampFieldName = "timestamp_custom"

	zerolog.DisableSampling(true)
	log.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

	if appName := viper.GetString("application.name"); appName != "" {
		log.Logger = log.Logger.With().Str("application", appName).Logger()
	}
}

func GetLogger() *zerolog.Logger {
	return &log.Logger
}
