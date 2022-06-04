package service

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

type NotifierService struct {
	db *sql.DB
}

func NewNotifierService(db *sql.DB) *NotifierService {
	return &NotifierService{db: db}
}

func (s *NotifierService) Run() {
	logFile := getLogFile("log.log")
	go s.notify(30*time.Second, logFile)
}

func (s *NotifierService) notify(duration time.Duration, file *os.File) {
	ticker := time.NewTicker(duration)

	for range ticker.C {
		slot24 := time.Now().Add(24 * time.Hour)
		slot2 := time.Now().Add(2 * time.Hour)

		farAppointments := s.checkDB(slot24)
		closeAppointments := s.checkDB(slot2)

		for _, appointment := range farAppointments {
			file.WriteString(fmt.Sprintf("%s | Привет %s! Напоминаем что вы записаны к %s завтра в %s!", time.Now().String(), appointment.UserName, appointment.DoctorSpec, slot24))
		}
		for _, appointment := range closeAppointments {
			file.WriteString(fmt.Sprintf("%s | Привет %s! Вам через 2 часа к %s в %s!", time.Now().String(), appointment.UserName, appointment.DoctorSpec, slot2))
		}
	}
}

func getLogFile(name string) *os.File {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	return f
}

type Notification struct {
	UserName   string
	DoctorSpec string
}

func (s *NotifierService) checkDB(timestamp time.Time) []Notification {
	query := fmt.Sprintf("SELECT doctors.spec, users.name " +
		"FROM doctors, users " +
		"WHERE doctors.id=(SELECT doctor_id FROM appointments WHERE slot=$1) " +
		"AND users.id=(SELECT user_id FROM appointments WHERE slot=$1)")
	rows, err := s.db.Query(query, timestamp)
	if err != nil {
		log.Printf("error while getting appointments with timestamp=%s", timestamp)
		return nil
	}

	var notification Notification
	var notifications []Notification
	for rows.Next() {
		if err := rows.Scan(&notification.DoctorSpec, &notification.UserName); err != nil {
			return nil
		}
		notifications = append(notifications, notification)
	}

	return notifications
}
