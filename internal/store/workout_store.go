package store

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func Open() (*sql.DB, error) {
	var (
		host     = "localhost"
		user     = "postgres"
		password = "postgres"
		dbname   = "postgres"
		port     = "5432"
		sslmode  = "disable"
	)

	dataSource := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)

	db, err := sql.Open("pgx", dataSource)

	if err != nil {
		return nil, fmt.Errorf("db: open %w", err)
	}

	fmt.Println("Connected to database...")

	return db, nil
}
