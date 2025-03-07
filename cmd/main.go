package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"platform-go-challenge/persistence"
	"platform-go-challenge/presentation"
	"time"

	"github.com/arangodb/go-driver"
	arangohttp "github.com/arangodb/go-driver/http"
	"github.com/gorilla/mux"
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
	serverOnPort := getEnv("SERVE_ON_PORT", "8000")

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

	r := mux.NewRouter()

	favouriteRoutes := presentation.CreateRoutes()
	for routePath, routeDefinition := range favouriteRoutes {
		log.Printf("adding %s route with methods %v\n", routePath, routeDefinition.Methods)
		r.Handle(routePath, routeDefinition.HandlerFunc).Methods(routeDefinition.Methods...)
	}

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+serverOnPort, r))
}
