package dbrepo

import (
	"errors"
	"github.com/techarm/go-bookings/internal/config"
	"github.com/techarm/go-bookings/internal/models"
	"time"
)

type testDBRepo struct {
	App *config.AppConfig
}

// InsertReservation 予約内容登録
func (m testDBRepo) InsertReservation(res models.Reservation) (int, error) {
	return 0, nil
}

// InsertRoomRestriction 予約制限内容登録
func (m testDBRepo) InsertRoomRestriction(rr models.RoomRestriction) error {
	return nil
}

// SearchRoomAvailabilityByDates 指摘期間内に部屋の空状況を取得
func (m testDBRepo) SearchRoomAvailabilityByDates(roomID int, start, end time.Time) (bool, error) {
	return true, nil
}

// SearchAvailabilityForAllRooms 現在全て予約可能の部屋情報を取得
func (m testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room
	return rooms, nil
}

// GetRoomById RoomIDでRoom情報を取得sるう
func (m testDBRepo) GetRoomById(id int) (models.Room, error) {
	var room models.Room
	if id == 100 {
		return room, errors.New("some error")
	}
	return room, nil
}
