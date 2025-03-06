package persistence

import (
	"log"

	"github.com/arangodb/go-driver"
)

func InitDB(client driver.Client, dbName string) error {
	clientRepository := NewClientRepository(client)
	_, err := clientRepository.GetOrCreateDB(dbName)
	if err != nil {
		log.Printf("%v", err)
		return err
	}

	return nil
}
