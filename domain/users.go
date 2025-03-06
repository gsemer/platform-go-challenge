package domain

import "time"

type User struct {
	Key       string    `json:"_key,omitempty"`
	FirstName string    `json:"first_name"`
	LatName   string    `json:"last_name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
