package repository

import (
	"clinic/dao"
	"database/sql"
	"fmt"
	"log"
)

type AppointmentsRepository struct {
	db *sql.DB
}

func NewAppointmentRepository(db *sql.DB) *AppointmentsRepository {
	return &AppointmentsRepository{db: db}
}

func (r *AppointmentsRepository) Create(dao dao.Appointment) error {
	query := fmt.Sprintf("INSERT INTO %s (doctor_id, user_id, slot) VALUES ($1, $2, $3)", appointmentsTable)
	_, err := r.db.Exec(query, dao.DoctorId, dao.UserId, dao.Slot)
	if err != nil {
		log.Printf("error while inserting new row in %s", appointmentsTable)
		return err
	}
	return nil
}
