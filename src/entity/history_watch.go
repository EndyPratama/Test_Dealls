package entity

import "time"

type HistoryWatch struct {
	ID        int64     `db:"id" json:"id"`
	Profile1  int64     `db:"profile1_id" json:"profile1_id"`
	Profile2  int64     `db:"profile2_id" json:"profile2_id"`
	Label     string    `db:"label" json:"label"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
