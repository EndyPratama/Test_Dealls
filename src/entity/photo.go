package entity

import "time"

type Photo struct {
	ID        int64     `db:"id" json:"id"`
	ProfileID int64     `db:"profile_id" json:"profile_id"`
	PhotoURL  string    `db:"photo_url" json:"photo_url"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
