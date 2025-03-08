package services

import "platform-go-challenge/domain"

type FavouriteService struct {
	fr domain.FavouriteRepository
}

func NewFavouriteService(fr domain.FavouriteRepository) *FavouriteService {
	return &FavouriteService{fr: fr}
}

func (fs FavouriteService) AddToFavourites(userID, assetID string) (domain.Favourite, error) {
	favourite, err := fs.fr.AddToFavourites(userID, assetID)
	if err != nil {
		return domain.Favourite{}, err
	}
	return favourite, nil
}
