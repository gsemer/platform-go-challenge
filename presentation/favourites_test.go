package presentation

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"platform-go-challenge/app/fakes"
	"platform-go-challenge/domain"
	"reflect"
	"testing"
	"time"

	"github.com/gorilla/mux"
)

func TestAddToFavourites(t *testing.T) {
	fs := &fakes.FakeFavouriteService{}

	favourite := domain.Favourite{
		From:      "user/1",
		To:        "asset/2",
		CreatedAt: time.Now(),
	}
	fs.AddToFavouritesReturns(favourite, nil)

	myHandler := FavouriteHandler{fs: fs}

	assetID := "2"
	request, _ := http.NewRequest("POST", "/assets/"+assetID+"/starred", nil)
	request.Header.Set("user_id", "1")

	response := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/assets/{asset_id}/starred", myHandler.AddToFavourites)
	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Status code is not %v, but is %v", http.StatusOK, response.Code)
		return
	}

	actualFavourite, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
		return
	}

	var result domain.Favourite
	err = json.Unmarshal(actualFavourite, &result)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(result.From, favourite.From) {
		t.Errorf("It was expected to have %v as a user id, but got %v instead", favourite.From, result.From)
		return
	}
	if !reflect.DeepEqual(result.To, favourite.To) {
		t.Errorf("It was expected to have %v as an asset id, but got %v instead", favourite.To, result.To)
		return
	}
}

func TestAddToFavourites_FAIL(t *testing.T) {
	fs := &fakes.FakeFavouriteService{}
	fs.AddToFavouritesReturns(domain.Favourite{}, errors.New("something went wrong"))

	myHandler := FavouriteHandler{fs: fs}

	assetID := "2"
	request, _ := http.NewRequest("POST", "/assets/"+assetID+"/starred", nil)
	request.Header.Set("user_id", "1")

	response := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/assets/{asset_id}/starred", myHandler.AddToFavourites)
	router.ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected %v but got %v", http.StatusBadRequest, response.Code)
	}
}
