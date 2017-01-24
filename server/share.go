package server

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/johnsudaar/moviie_api/models"
)

func ShareMovie(req *http.Request, params martini.Params) (int, string) {
	apiKey := req.URL.Query().Get("api_key")
	friendKey := params["friend"]

	me, err := models.FindUserByApiKey(apiKey)

	if err != nil {
		return 500, err.Error()
	}

	if me == nil {
		return 400, "Invalid API KEY"
	}

	friend, err := models.FindUserByUsername(friendKey)

	if err != nil {
		return 500, err.Error()
	}

	if friend == nil {
		return 400, "Unknown friend"
	}

	correct := false
	for _, f := range me.Friends {
		if f == friend.Username {
			correct = true
			break
		}
	}

	if !correct {
		return 400, "You are not friend with this person"
	}

	var m models.Movie

	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&m)

	if err != nil {
		return 400, "Invalid movie"
	}

	s := models.Share{
		From:    me.Username,
		To:      friend.Username,
		Movie:   m,
		MovieID: m.ID,
	}

	err = s.Save()

	if err != nil {
		return 500, err.Error()
	}

	return 200, "Ok"
}

func MyShares(req *http.Request) (int, string) {
	apiKey := req.URL.Query().Get("api_key")

	u, err := models.FindUserByApiKey(apiKey)

	if err != nil {
		return 500, err.Error()
	}

	if u == nil {
		return 400, "Invalid API KEY"
	}

	shares, err := models.FindShareByUser(u.Username)

	if err != nil && err.Error() != "not found" {
		return 500, err.Error()
	}

	if shares == nil {
		shares = make([]models.Share, 0)
	}
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	err = encoder.Encode(shares)

	if err != nil {
		return 500, err.Error()
	}

	return 200, buffer.String()
}
