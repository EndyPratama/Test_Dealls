package entity

import "time"

type Likes struct {
	ID             int64     `db:"id" json:"id"`
	HistoryWatchID int64     `db:"history_watch_id" json:"history_watch_id"`
	MatchesID      int64     `db:"matches_id" json:"matches_id"`
	LikerID        int64     `db:"liker_id" json:"liker_id"`
	LikerName      string    `db:"liker_name" json:"liker_name"`
	LikedID        int64     `db:"liked_id" json:"liked_id"`
	LikedName      string    `db:"liked_name" json:"liked_name"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
}
