package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/rennanbadaro/discount-calculator/infrastructure/config"
)

type PostgresClient struct {
	Conn *sql.DB
}

func NewPostgresClient() (*PostgresClient, error) {
	dbConfig := config.Config().PostgresConfig

	pgConnString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
	)

	conn, err := sql.Open("postgres", pgConnString)

	if err != nil {
		fmt.Println(err)

		return nil, err
	}

	err = conn.Ping()

	if err != nil {
		fmt.Println(err)

		return nil, err
	}

	return &PostgresClient{Conn: conn}, nil
}
