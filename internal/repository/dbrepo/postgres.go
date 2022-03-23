package dbrepo

import (
	"context"
	"database/sql"
	"github.com/techarm/go-bookings/internal/config"
	"github.com/techarm/go-bookings/internal/models"
	"time"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

// InsertReservation 予約内容登録
func (m postgresDBRepo) InsertReservation(res models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into reservations (first_name, last_name, phone, email,
             start_date, end_date, room_id, created_at, updated_at)
             values ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id`

	var id int
	row := m.DB.QueryRowContext(ctx, stmt,
		res.FirstName, res.LastName, res.Phone, res.Email,
		res.StartDate, res.EndDate, res.RoomID, time.Now(), time.Now())
	err := row.Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

// InsertRoomRestriction 予約制限内容登録
func (m postgresDBRepo) InsertRoomRestriction(rr models.RoomRestriction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into room_restrictions(start_date, end_date, room_id,
             reservation_id, restriction_id, created_at, updated_at)
             values($1, $2, $3, $4, $5, $6, $7)`

	_, err := m.DB.ExecContext(ctx, stmt, rr.StartDate,
		rr.EndDate, rr.RoomID, rr.ReservationID, rr.RestrictionID, time.Now(), time.Now())

	if err != nil {
		return err
	}

	return nil
}
