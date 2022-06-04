package service

import (
	"clinic/internal/repository"
	"clinic/model"
)

type AppointmentsService struct {
	appointmentRepo repository.Appointments
}

func NewAppointmentService(appointmentsRepo repository.Appointments) *AppointmentsService {
	return &AppointmentsService{
		appointmentRepo: appointmentsRepo,
	}
}

func (s *AppointmentsService) Add(appointment model.Appointment) error {
	dao := ToAppointmentDao(appointment)

	err := s.appointmentRepo.Create(dao)
	if err != nil {
		return err
	}

	return nil
}
