package database

import (
	"context"
	"errors"
	"os"

	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

func Connect() error {
	var err error
	conn, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URI"))
	if err == nil {
		var ping int
		conn.QueryRow(context.Background(), "SELECT 1 AS ping").Scan(&ping)

		if ping != 1 {
			err = errors.New("failed to ping database")
		}
	}
	return err
}

func Get() *pgx.Conn {
	return conn
}
