package service

import (
	"clinic/internal/repository"
	"clinic/model"
)

type Appointment interface {
	Add(appointment model.Appointment) error
}

type Services struct {
	Appointment
}

func NewServices(repos *repository.Repositories) *Services {
	return &Services{
		Appointment: NewAppointmentService(repos.Appointments),
	}
}
