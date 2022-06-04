package service

import (
	"clinic/dao"
	"clinic/model"
)

func ToAppointmentDao(appointment model.Appointment) dao.Appointment {
	return dao.Appointment{
		DoctorId: appointment.DoctorId,
		UserId:   appointment.UserId,
		Slot:     appointment.Slot,
	}
}
