package server

import (
	"net/http"

	"github.com/go-martini/martini"
	"github.com/johnsudaar/moviie_api/models"
)

func AddFriend(req *http.Request, params martini.Params) (int, string) {
	friendName := params["friend"]
	apiKey := req.URL.Query().Get("api_key")

	me, err := models.FindUserByApiKey(apiKey)

	if err != nil {
		return 500, err.Error()
	}

	friend, err := models.FindUserByUsername(friendName)
	if err != nil {
		return 500, err.Error()
	}

	if me == nil {
		return 400, "Invalid API Key"
	}

	if friend == nil {
		return 400, "Friend not found"
	}

	err = me.AddFriend(friend)
	if err != nil {
		return 500, err.Error()
	}
	return 200, "Ok"
}
