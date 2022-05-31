package service

import (
	"clinic/internal/repository"
	"database/sql"
)

type Services struct {
}

func NewServices(repos *repository.Repositories, db *sql.DB) *Services {
	return &Services{}
}
