package main

import (
	"context"
	"log"
	"time"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

func main() {
	log.Println("Server is up and running...")

	log.Println("Waiting 10 seconds for the ArangoDB container to be ready...")
	time.Sleep(10 * time.Second)

	connection, err := http.NewConnection(
		http.ConnectionConfig{
			Endpoints: []string{"http://arangodb:8529"},
		},
	)
	if err != nil {
		log.Fatalf("Failed to create HTTP connection: %v", err)
	}

	client, err := driver.NewClient(
		driver.ClientConfig{
			Connection:     connection,
			Authentication: driver.BasicAuthentication("root", "rootpassword"),
		},
	)
	if err != nil {
		log.Fatalf("Failed to create the client: %v", err)
	}

	_, err = client.Database(context.Background(), "platform-db")
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
