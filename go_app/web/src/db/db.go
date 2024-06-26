package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Settings struct {
	Host     string
	Port     int
	User     string
	Password string
	DBname   string
}

func Connect(s *Settings) (*sql.DB, error) {
	//create connection string
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", s.Host, s.Port, s.User, s.Password, s.DBname)

	// open database connection
	return sql.Open("postgres", connStr)
}
