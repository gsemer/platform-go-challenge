package presentation

import (
	"encoding/json"
	"net/http"
	"platform-go-challenge/domain"

	"github.com/gorilla/mux"
)

type FavouriteHandler struct {
	fs domain.FavouriteService
}

func NewFavouriteHandler(fs domain.FavouriteService) *FavouriteHandler {
	return &FavouriteHandler{fs: fs}
}

func (fh FavouriteHandler) AddToFavourites(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	assetID, ok := vars["asset_id"]
	if !ok {
		writer.WriteHeader(http.StatusBadRequest)
		result, _ := json.Marshal("Asset ID is missing")
		writer.Write(result)
		return
	}

	userID := request.Header.Get("user_id")

	favourite, err := fh.fs.AddToFavourites(userID, assetID)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		result, _ := json.Marshal(err)
		writer.Write(result)
		return
	}

	writer.WriteHeader(http.StatusOK)
	result, _ := json.Marshal(favourite)
	writer.Write(result)
}
