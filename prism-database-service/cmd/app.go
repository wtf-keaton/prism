package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

func main() {
	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URI"))
	if err != nil {
		log.Fatalf("Error: failed to connect to database: %v", err)
	}

	defer func() {
		if err := conn.Close(ctx); err != nil {
			log.Printf("Warning: error while closing the database connection: %v", err)
		} else {
			log.Println("Database connection closed successfully")
		}
	}()

	if err := pingDatabase(ctx, conn); err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Println("Successfully connected to the database")

	sqlFilePath := "db_1.sql"
	data, err := readSQLFile(sqlFilePath)
	if err != nil {
		log.Fatalf("Error reading SQL file '%s': %v", sqlFilePath, err)
	}

	if err := executeMigration(ctx, conn, string(data)); err != nil {
		log.Fatalf("Error executing migration: %v", err)
	}

	log.Println("Database migration successfully completed")
}

func pingDatabase(ctx context.Context, conn *pgx.Conn) error {
	var ping int
	err := conn.QueryRow(ctx, "SELECT 1 AS ping").Scan(&ping)
	if err != nil || ping != 1 {
		return fmt.Errorf("failed to ping database: %v", err)
	}
	return nil
}

func readSQLFile(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read SQL file '%s': %v", filePath, err)
	}
	return data, nil
}

func executeMigration(ctx context.Context, conn *pgx.Conn, sql string) error {
	_, err := conn.Exec(ctx, sql)
	if err != nil {
		return fmt.Errorf("failed to execute migration: %v", err)
	}
	return nil
}
