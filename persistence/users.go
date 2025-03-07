package persistence

import (
	"context"
	"errors"
	"log"
	"platform-go-challenge/consts"
	"platform-go-challenge/domain"

	"github.com/arangodb/go-driver"
)

type UserRepository struct {
	db driver.Database
}

func NewUserRepository(db driver.Database) *UserRepository {
	return &UserRepository{db: db}
}

func (ur UserRepository) CreateUsers() ([]domain.User, error) {
	userCollection, err := ur.db.Collection(context.Background(), "user")
	if err != nil {
		log.Println("Couldn't get the collection when trying to create users")
	}

	numOfDocs, err := userCollection.Count(context.Background())
	if err != nil {
		log.Println("Failed to get document count for user collection")
	}
	if numOfDocs > 0 {
		return []domain.User{}, errors.New("there are already documents in user collections")
	}

	users := consts.Users

	arangoUsers := []domain.User{}
	for _, user := range users {
		meta, err := userCollection.CreateDocument(context.Background(), &user)
		if err != nil {
			log.Printf("Couldn't create new document")
			return []domain.User{}, err
		}
		user.Key = meta.Key
		arangoUsers = append(arangoUsers, user)
	}
	return arangoUsers, nil
}
