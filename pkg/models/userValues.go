package models

import (
	"context"
	"database/sql"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

var (
	DB     *sql.DB
	Err    error
	Rdb    *redis.Client
	Ctx    = context.Background()
	Log    = logrus.New()
	Jwtkey = []byte("a-string-secret-at-least-256-bits-long")
)
