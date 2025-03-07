package domain

import "time"

type Favourite struct {
	From      string    `json:"_from"` // Reference to the user document
	To        string    `json:"_to"`   // Reference to the asset document
	CreatedAt time.Time `json:"created_at"`
}

type FavouriteService interface {
}

type FavouriteRepository interface {
}
