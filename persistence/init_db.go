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
	dbExists, err := cr.client.DatabaseExists(context.Background(), dbName)
	if err != nil {
		log.Printf("Failed to check if database exists")
	}
	var db driver.Database
	if !dbExists {
		log.Printf("Database doen't exist")
		db, err = cr.client.CreateDatabase(context.Background(), dbName, nil)
		if err != nil {
			log.Printf("Failed to create database")
			return nil, err
		}
	} else {
		log.Printf("Database already exists")
		db, err = cr.client.Database(context.Background(), dbName)
		if err != nil {
			log.Printf("Failed to open database")
			return nil, err
		}
	}
	return db, nil
}
