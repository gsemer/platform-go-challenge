package presentation

import "platform-go-challenge/domain"

func CreateRoutes(fs domain.FavouriteService) map[string]domain.RouteDefinition {
	fh := NewFavouriteHandler(fs)

	return map[string]domain.RouteDefinition{
		"/assets/{asset_id}/starred": {
			Methods:     []string{"POST"},
			HandlerFunc: fh.AddToFavourites,
		},
	}
}
