package server

import (
	"encoding/json"
	"net/http"

	"github.com/johnsudaar/moviie_api/models"
)

func SetMovies(req *http.Request) (int, string) {
	apiKey := req.URL.Query().Get("api_key")
	user, err := models.FindUserByApiKey(apiKey)

	if err != nil {
		return 500, err.Error()
	}

	if user == nil {
		return 400, "Invalid API KEY"
	}

	var movies []*models.Movie

	decoder := json.NewDecoder(req.Body)

	err = decoder.Decode(&movies)

	if err != nil {
		return 500, err.Error()
	}

	user.Movies = movies

	err = user.Update()

	if err != nil {
		return 500, err.Error()
	}

	return 200, "Ok"
}
