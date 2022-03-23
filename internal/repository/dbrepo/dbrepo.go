package dbrepo

import (
	"database/sql"
	"github.com/techarm/go-bookings/internal/config"
	"github.com/techarm/go-bookings/internal/repository"
)

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
