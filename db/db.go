package db

import "database/sql"

type PostgresDB struct {
	DB *sql.DB
}

func NewPostgresDB(dataSourceName string) (*PostgresDB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	//Ensure the database is reachable
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresDB{DB: db}, nil
}
