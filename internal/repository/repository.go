package repository

import (
	"github.com/techarm/go-bookings/internal/models"
	"time"
)

type DatabaseRepo interface {
	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(res models.RoomRestriction) error
	SearchRoomAvailabilityByDates(roomId int, start, end time.Time) (int, error)
	SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error)
}
