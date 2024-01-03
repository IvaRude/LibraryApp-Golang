package db

import (
	"context"
	"fmt"

	"homework-3/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewDB(ctx context.Context, config *config.Config) (*Database, error) {
	pool, err := pgxpool.Connect(ctx, generateDsn(config))
	if err != nil {
		return nil, err
	}
	return newDatabase(pool), nil
}

func generateDsn(config *config.Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.DbName)
}
