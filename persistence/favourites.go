package persistence

import (
	"context"
	"log"
	"platform-go-challenge/domain"
	"time"

	"github.com/arangodb/go-driver"
)

type FavouriteRepository struct {
	db         driver.Database
	collection map[string]driver.Collection
}

func NewFavouriteRepository(db driver.Database, collection map[string]driver.Collection) *FavouriteRepository {
	return &FavouriteRepository{db: db, collection: collection}
}

func (fr FavouriteRepository) AddToFavourites(userID, assetID string) (domain.Favourite, error) {
	favourite := domain.Favourite{
		From:      userID,
		To:        assetID,
		CreatedAt: time.Now(),
	}
	_, err := fr.collection["favourite"].CreateDocument(context.Background(), &favourite)
	if err != nil {
		log.Println("Couldn't create new document to favourites")
		return domain.Favourite{}, err
	}
	return favourite, nil
}
