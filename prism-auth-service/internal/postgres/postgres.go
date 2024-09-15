package postgres

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

var (
	conn    *pgx.Conn
	timeout = 5 * time.Second
)

var (
	ErrUserExists   = errors.New("user already exists")
	ErrUserNotFound = errors.New("user not found")
)

func Connect() error {
	if conn != nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var err error
	conn, err = pgx.Connect(ctx, os.Getenv("DATABASE_URI"))
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	var ping int
	err = conn.QueryRow(ctx, "SELECT 1 AS ping").Scan(&ping)
	if err != nil || ping != 1 {
		conn.Close(ctx)
		return errors.New("failed to ping database")
	}

	log.Println("Successfully connected to the database")
	return err
}

func Get() *pgx.Conn {
	if conn == nil {
		log.Println("Warning: Attempt to get connection before initializing")
	}

	return conn
}

func Close() error {
	if conn != nil {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		err := conn.Close(ctx)
		if err != nil {
			return fmt.Errorf("failed to close database connection: %v", err)
		}

		log.Println("Database connection closed successfully")
		conn = nil
	}
	return nil
}
