package repository

import (
	"clinic/dao"
	"database/sql"
)

type Appointments interface {
	Create(dao dao.Appointment) error
}

type Repositories struct {
	Appointments
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Appointments: NewAppointmentRepository(db),
	}
}
