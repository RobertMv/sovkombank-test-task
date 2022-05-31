package repository

import "database/sql"

type Repositories struct {
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{}
}
