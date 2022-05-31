package dao

import "time"

type User struct {
	Id      int64
	Phone   string
	Name    string
	Surname string
}

type Doctor struct {
	Id      int64
	Name    string
	Surname string
}

type Slots struct {
	Id       int64
	DoctorId int64
	Slot     time.Time
}

type Appointment struct {
	DoctorId int64
	UserId   int64
	Slot     time.Time
}
