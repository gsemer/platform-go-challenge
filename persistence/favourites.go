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
	user, err := fr.collection["user"].ReadDocument(context.Background(), userID, nil)
	if err != nil {
		return domain.Favourite{}, err
	}

	asset, err := fr.collection["asset"].ReadDocument(context.Background(), assetID, nil)
	if err != nil {
		return domain.Favourite{}, err
	}

	favourite := domain.Favourite{
		From:      string(user.ID),
		To:        string(asset.ID),
		CreatedAt: time.Now(),
	}
	_, err = fr.collection["favourite"].CreateDocument(context.Background(), &favourite)
	if err != nil {
		log.Println("Couldn't create new document to favourites")
		return domain.Favourite{}, err
	}
	return favourite, nil
}
