package presentation

import "platform-go-challenge/domain"

func CreateRoutes() map[string]domain.RouteDefinition {
	_ = NewFavouriteHandler()

	return map[string]domain.RouteDefinition{}
}
