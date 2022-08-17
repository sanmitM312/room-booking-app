package dbrepo

import (
	"database/sql"

	"github.com/sanmitM312/room-booking-app/internal/config"
	"github.com/sanmitM312/room-booking-app/internal/repository"
)

type postgresDBRepo struct{
	App *config.AppConfig
	DB *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo{
	return &postgresDBRepo{
		App : a,
		DB: conn,
	}
}