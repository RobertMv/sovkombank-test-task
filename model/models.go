package model

import "time"

type User struct {
	Id      string `json:"id"`
	Phone   string `json:"phone"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type Doctor struct {
	Id    string      `json:"id"`
	Name  string      `json:"name"`
	Spec  string      `json:"spec"`
	Slots []time.Time `json:"slots"`
}

type Appointment struct {
	DoctorId string    `json:"doctorId"`
	UserId   string    `json:"userId"`
	Slot     time.Time `json:"slot"`
}
