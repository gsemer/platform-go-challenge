package persistence

import (
	"context"
	"log"

	"github.com/arangodb/go-driver"
)

type ClientRepository struct {
	client driver.Client
}

func NewClientRepository(client driver.Client) *ClientRepository {
	return &ClientRepository{client: client}
}

func (cr ClientRepository) GetOrCreateDB(dbName string) (driver.Database, error) {
	db_exists, err := cr.client.DatabaseExists(context.Background(), dbName)
	if err != nil {
		log.Printf("Failed to check if database exists")

	}
	var db driver.Database
	if !db_exists {
		db, err = cr.client.CreateDatabase(context.Background(), dbName, nil)
		if err != nil {
			log.Printf("Failed to create database")
			return nil, err
		}
	} else {
		db, err = cr.client.Database(context.Background(), dbName)
		if err != nil {
			log.Printf("Failed to open database")
			return nil, err
		}
	}
	return db, nil
}
