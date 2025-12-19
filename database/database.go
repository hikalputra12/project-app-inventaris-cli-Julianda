package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

/*
--- InitDB initializes and returns a connection to the PostgreSQL database using the pgx driver. ---
*/
func InitDB() (*pgx.Conn, error) {
	ConnStr := "user=postgres password=Putra1014 dbname=inventory host=localhost port=5432 sslmode=disable"
	conn, err := pgx.Connect(context.Background(), ConnStr)
	return conn, err
}
