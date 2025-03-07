package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
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

	connection, err := http.NewConnection(
		http.ConnectionConfig{
			Endpoints: []string{getEnv("ARANGO_URL", "")},
		},
	)
	if err != nil {
		log.Fatalf("Failed to create HTTP connection: %v", err)
	}

	client, err := driver.NewClient(
		driver.ClientConfig{
			Connection: connection,
			Authentication: driver.BasicAuthentication(
				getEnv("ARANGO_USERNAME", "root"),
				getEnv("ARANGO_PASSWORD", ""),
			),
		},
	)
	if err != nil {
		log.Fatalf("Failed to create the client: %v", err)
	}

	_, err = client.Database(context.Background(), getEnv("ARANGO_DB", ""))
	if err != nil {
		if driver.IsNotFoundGeneral(err) {
			_, err = client.CreateDatabase(context.Background(), "platform-db", nil)
			if err != nil {
				log.Fatalf("Failed to create database: %v", err)
			}
		}
		log.Println("Database 'platform-db' was created successfully")
	}
}
