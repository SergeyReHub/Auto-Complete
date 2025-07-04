package postgres_storage

import (
	"auto_complite/internal/config"
	pool_conns "auto_complite/internal/repository/postgres_storage/pool_connections"
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type PostgresStorage struct {
	Pool   *pgxpool.Pool
	Logger *zap.Logger
}

func PostgresInit(cfg *config.DB, logger *zap.Logger) (PostgresUsecase, error) {
	pool, err := pool_conns.CreatePool(cfg)
	if err != nil {
		return nil, errors.New("DB postgres error. Init pool error.\n" + err.Error())
	}
	return &PostgresStorage{
		Pool:   pool,
		Logger: logger,
	}, nil
}

func (p *PostgresStorage) GetSimilarsFromDb(ctx context.Context, str string) ([]string, error) {
	conn, err := p.Pool.Acquire(ctx)
	if err != nil {
		p.Logger.Error("DB Postgres error. Create connection error.", zap.Error(err))
		return nil, errors.New("DB Postgres error. Create connection error. " + err.Error())
	}
	defer conn.Release()

	rows, err := conn.Query(ctx, "SELECT phrase FROM list_default_phrases WHERE phrase ILIKE "+"'"+str+"%'")
	defer rows.Close()
	if err != nil {
		p.Logger.Error("DB Postgres error. Query GetSimilars error.", zap.String("string to find similars", str))
		return nil, errors.New("DB Postgres error. Query GetSimilars error. " + err.Error())
	}
	var phrases []string
	for rows.Next() {
		var phrase string
		err = rows.Scan(&phrase)
		if err != nil {
			p.Logger.Error("DB Postgres error. Row scan error.", zap.Error(err))
			return nil, errors.New("DB Postgres error. Row scan error. " + err.Error())
		}
		phrases = append(phrases, phrase)
	}
	return phrases, nil
}
