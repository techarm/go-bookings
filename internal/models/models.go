package models

import "time"

// User ユーザー情報モデル
type User struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreateAt    time.Time
	UpdateAt    time.Time
}

// Room roomモデル
type Room struct {
	ID       int
	RoomName string
	CreateAt time.Time
	UpdateAt time.Time
}

// Restriction 予約制限モデル
type Restriction struct {
	ID              int
	RestrictionName string
	CreateAt        time.Time
	UpdateAt        time.Time
}

// Reservation 予約モデル
type Reservation struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	StartDate time.Time
	EndDate   time.Time
	RoomID    int
	CreateAt  time.Time
	UpdateAt  time.Time
	Room      Room
}

// RoomRestriction 予約room制限モデル
type RoomRestriction struct {
	ID            int
	StartDate     time.Time
	EndDate       time.Time
	RoomID        int
	ReservationID int
	RestrictionID int
	CreateAt      time.Time
	UpdateAt      time.Time
	Room          Room
	Reservation   Reservation
	Restriction   Restriction
}
