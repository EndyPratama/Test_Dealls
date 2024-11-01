package entity

import "time"

type User struct {
	ID             int64     `db:"id" json:"id"`
	Email          string    `db:"email" json:"email"`
	Password       string    `db:"password" json:"password"`
	SubscriptionID int64     `db:"subscription_id" json:"subscription_id"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
}
