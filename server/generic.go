package server

import "net/http"

func ServerError(err string, res http.ResponseWriter) {
	res.WriteHeader(500)
	res.Write([]byte(err))
}

func UserError(err string, res http.ResponseWriter) {
	res.WriteHeader(400)
	res.Write([]byte(err))
}

func Ok(val string, res http.ResponseWriter) {
	res.WriteHeader(200)
	res.Write([]byte(val))
}
