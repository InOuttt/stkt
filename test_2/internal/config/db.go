package config

import (
	"context"
	"time"
)

type DbConfig struct {
	DSN     string
	Timeout time.Duration
}

func GetDbConf(args Args) DbConfig {
	return DbConfig{
		DSN:     args.MysqlDsn,
		Timeout: args.DbTimeout,
	}
}

func NewCtx(duration time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), duration)
}
