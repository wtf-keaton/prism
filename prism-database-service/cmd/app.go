package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URI"))
	if err == nil {
		var ping int
		conn.QueryRow(context.Background(), "SELECT 1 AS ping").Scan(&ping)

		if ping != 1 {
			log.Fatalf("Failed to ping database")
		}
	}

}
