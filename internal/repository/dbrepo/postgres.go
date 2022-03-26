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

// SearchRoomAvailabilityByDates 指摘期間内に部屋の空状況を取得
func (m postgresDBRepo) SearchRoomAvailabilityByDates(roomID int, start, end time.Time) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `select count(*) from room_restrictions where room_id = $1 and $2 < end_date and $3 < start_date`
	row := m.DB.QueryRowContext(ctx, stmt, roomID, start, end)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}

	return count == 0, nil
}

// SearchAvailabilityForAllRooms 現在全て予約可能の部屋情報を取得
func (m postgresDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
		select
			r.id,
			r.room_name
		from
			rooms r
		where
			r.id not in (
				select
					rr.room_id
				from
					room_restrictions rr
				where
					$1 < end_date
					and $2 > start_date
			)`

	var rooms []models.Room
	rows, err := m.DB.QueryContext(ctx, stmt, start, end)
	if err != nil {
		return rooms, err
	}

	for rows.Next() {
		var room models.Room
		err = rows.Scan(&room.ID, &room.RoomName)
		if err != nil {
			return rooms, err
		}
		rooms = append(rooms, room)
	}

	if err = rows.Err(); err != nil {
		return rooms, err
	}

	return rooms, nil
}

// GetRoomById RoomIDでRoom情報を取得sるう
func (m postgresDBRepo) GetRoomById(id int) (models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `select id, room_name, created_at, updated_at from rooms where id = $1`
	row := m.DB.QueryRowContext(ctx, stmt, id)

	var room models.Room
	err := row.Scan(&room.ID, &room.RoomName, &room.UpdatedAt, &room.CreatedAt)
	if err != nil {
		return room, err
	}

	return room, nil
}
