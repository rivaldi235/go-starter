package config

import (
	"database/sql"
	"fmt"
	"service-code/model/dto"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
)

func ConnectToDB(in dto.ConfigData, logger zerolog.Logger) (*sql.DB, error) {
	// code connect to db
	logger.Info().Msg("Trying to conncet DB . . .")

	//initialize database conncetion
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", in.DbConfig.Host, in.DbConfig.User, in.DbConfig.Pass, in.DbConfig.Database, in.DbConfig.DbPort)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to open database conncetion")
		return nil, err
	}

	logger.Info().Msg("Successfully conncected to the database")
	return db, nil
}
