package main

import (
	"context"
	"log"
	"os"
	"platform-go-challenge/persistence"
	"time"

	"github.com/arangodb/go-driver"
	arangohttp "github.com/arangodb/go-driver/http"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	log.Println("Server is up and running...")

	log.Println("Waiting 10 seconds for the ArangoDB container to be ready...")
	time.Sleep(10 * time.Second)

	arangoURL := getEnv("ARANGO_URL", "")
	arangoUsername := getEnv("ARANGO_USERNAME", "root")
	arangoPassword := getEnv("ARANGO_PASSWORD", "")
	arangoDB := getEnv("ARANGO_DB", "")

	connection, err := arangohttp.NewConnection(
		arangohttp.ConnectionConfig{
			Endpoints: []string{arangoURL},
		},
	)
	if err != nil {
		log.Fatalf("Failed to create HTTP connection: %v", err)
	}

	client, err := driver.NewClient(
		driver.ClientConfig{
			Connection:     connection,
			Authentication: driver.BasicAuthentication(arangoUsername, arangoPassword),
		},
	)
	if err != nil {
		log.Fatalf("Failed to create the client: %v", err)
	}

	db, err := client.Database(context.Background(), arangoDB)
	if err != nil {
		log.Printf("Error getting database %s: %v", arangoDB, err)
		if driver.IsNotFoundGeneral(err) {
			db, err = client.CreateDatabase(context.Background(), arangoDB, nil)
			if err != nil {
				log.Fatalf("Failed to create database: %v", err)
			}
		}
		log.Println("Database 'platform-db' was created successfully")
	}

	collectionNames := []string{"user", "asset", "favourite"}
	for _, collectionName := range collectionNames {
		_, err := db.Collection(context.Background(), collectionName)
		if err != nil {
			if driver.IsNotFoundGeneral(err) {
				_, err = db.CreateCollection(context.Background(), collectionName, nil)
				if err != nil {
					log.Fatalf("Failed to create collection: %v", err)
				}
			}
			log.Printf("Collection %s was created successfully", collectionName)
		}
	}

	userRepo := persistence.NewUserRepository(db)
	assetRepo := persistence.NewAssetRepository(db)

	_, _ = userRepo.CreateUsers()
	_, _ = assetRepo.CreateAssets()
}
