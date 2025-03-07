package persistence

import (
	"context"
	"errors"
	"log"
	"platform-go-challenge/consts"
	"platform-go-challenge/domain"

	"github.com/arangodb/go-driver"
)

type AssetRepository struct {
	db driver.Database
}

func NewAssetRepository(db driver.Database) *AssetRepository {
	return &AssetRepository{db: db}
}

func (ar AssetRepository) CreateAssets() ([]domain.Asset, error) {
	assetCollection, err := ar.db.Collection(context.Background(), "asset")
	if err != nil {
		log.Println("Couldn't get the collection when trying to create assets")
	}

	numOfDocs, err := assetCollection.Count(context.Background())
	if err != nil {
		log.Println("Failed to get document count for asset collection")
	}
	if numOfDocs > 0 {
		return []domain.Asset{}, errors.New("there are already documents in asset collections")
	}

	assets := append(append(consts.Charts, consts.Insights...), consts.Audience...)

	arangoAssets := []domain.Asset{}
	for _, asset := range assets {
		meta, err := assetCollection.CreateDocument(context.Background(), &asset)
		if err != nil {
			log.Printf("Couldn't create new document")
			return []domain.Asset{}, err
		}
		asset.Key = meta.Key
		arangoAssets = append(arangoAssets, asset)
	}
	return arangoAssets, nil
}
