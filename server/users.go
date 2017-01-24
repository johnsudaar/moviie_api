package server

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/johnsudaar/moviie_api/models"
	"github.com/johnsudaar/moviie_api/requests"
)

func Login(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var login requests.Login
	err := decoder.Decode(&login)

	if err != nil {
		UserError("Invalid data", res)
		return
	}

	user, err := models.Login(login.Username, login.Password)

	if err != nil {
		UserError(err.Error(), res)
		return
	}

	res.WriteHeader(200)
	encoder := json.NewEncoder(res)
	encoder.Encode(user)
}

func Register(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var register requests.Login
	err := decoder.Decode(&register)

	if err != nil {
		UserError("Invalid data", res)
		return
	}

	user := &models.User{
		Username: register.Username,
		Password: register.Password,
	}

	err = user.Save()

	if err != nil {
		ServerError(err.Error(), res)
		return
	}

	Ok("Ok", res)
}

func FindUser(params martini.Params) (int, string) {
	username := params["user"]
	user, err := models.FindUserByUsername(username)
	if err != nil {
		return 500, err.Error()

	}
	var res bytes.Buffer
	encoder := json.NewEncoder(&res)
	user.ApiKey = "[REDACTED]"
	encoder.Encode(user)
	return 200, res.String()
}

func SearchUser(r *http.Request) (int, string) {
	pattern := r.URL.Query().Get("pattern")
	users, err := models.SearchUser(pattern)
	if err != nil {
		return 500, err.Error()
	}

	var res bytes.Buffer
	encoder := json.NewEncoder(&res)
	encoder.Encode(users)
	return 200, res.String()
}
