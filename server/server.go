package server

import (
	"github.com/go-martini/martini"
	"github.com/johnsudaar/moviie_api/configuration"
)

func Launch() {
	m := martini.Classic()

	m.Post("/users/login", Login)
	m.Post("/users", Register)
	m.Get("/users/:user", FindUser)
	m.Post("/users/:friend", AddFriend)
	m.Post("/user/movies", SetMovies)
	m.Get("/user/search", SearchUser)
	m.RunOnAddr(":" + configuration.E["PORT"])
}
