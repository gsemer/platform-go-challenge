package main

import (
	"fmt"
	"log"
	"platform-go-challenge/persistence"
	"time"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Server is up and running...")
		time.Sleep(2 * time.Second)
	}

	connection, err := http.NewConnection(
		http.ConnectionConfig{
			Endpoints: []string{"http://arangodb_go_container:8529"},
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

	clientRepository := persistence.NewClientRepository(client)
	_, err = clientRepository.GetOrCreateDB("favourite")
	if err != nil {
		log.Fatalf("%v", err)
	}
}
