package persistence

import (
	"context"
	"log"

	"github.com/arangodb/go-driver"
)

type CollectionRepository struct {
	db driver.Database
}

func NewCollectionRepository(db driver.Database) *CollectionRepository {
	return &CollectionRepository{db: db}
}

func (cr CollectionRepository) GetOrCreate(name string) driver.Collection {
	collection, err := cr.db.Collection(context.Background(), name)
	if err != nil {
		if driver.IsNotFoundGeneral(err) {
			collection, err = cr.db.CreateCollection(context.Background(), name, nil)
			if err != nil {
				log.Fatalf("Failed to create collection: %v", err)
			}
		}
		log.Printf("Collection %s was created successfully", name)
	}
	return collection
}
