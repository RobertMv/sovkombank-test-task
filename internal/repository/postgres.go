package repository

import "database/sql"

const (
	usersTable        = "users"
	doctorsTable      = "doctors"
	slotsTable        = "slots"
	appointmentsTable = "appointments"
)

func NewPostgresDB(dbUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
