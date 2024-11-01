package entity

import "time"

type Profile struct {
	ID             int64     `db:"id" json:"id"`
	UserID         int64     `db:"user_id" json:"user_id"`
	Name           string    `db:"name" json:"name"`
	Gender         string    `db:"gender" json:"gender"`
	Bio            string    `db:"bio" json:"bio"`
	Birthdate      string    `db:"birthdate" json:"birthdate"`
	Location       string    `db:"location" json:"location"`
	SubscriptionID int64     `db:"subscription_id" json:"subscription_id"`
	HistoryWatchID int64     `db:"history_watch_id" json:"history_watch_id"`
	Label          string    `db:"label" json:"label"`
	Photo          []string  `db:"photo" json:"photo"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
}
