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
	m.Get("/user", FindByApiKey)
	m.Post("/users/:friend", AddFriend)
	m.Post("/user/movies", SetMovies)
	m.Get("/user/search", SearchUser)
	m.Post("/share/:friend", ShareMovie)
	m.Get("/share", MyShares)
	m.RunOnAddr(":" + configuration.E["PORT"])
}
