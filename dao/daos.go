package dao

import "time"

type User struct {
	Id      string
	Phone   string
	Name    string
	Surname string
}

type Doctor struct {
	Id   string
	Name string
	Spec string
}

type Slots struct {
	DoctorId string
	Slot     time.Time
}

type Appointment struct {
	DoctorId string
	UserId   string
	Slot     time.Time
}
