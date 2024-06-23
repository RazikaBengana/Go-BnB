package repository

import "github.com/RazikaBengana/Go-BnB/internal/models"

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) error
}
