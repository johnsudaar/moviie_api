package configuration

import (
	"log"
	"os"

	"gopkg.in/mgo.v2"
)

var E = map[string]string{
	"MONGO_URL": "mongodb://localhost:27017/moviie",
	"MONGO_DB":  "moviie",
	"PORT":      "8080",
}

var MongoSession *mgo.Session

func init() {
	for name := range E {
		envValue := os.Getenv(name)
		if envValue != "" {
			E[name] = envValue
		}
	}

	session, err := mgo.Dial(E["MONGO_URL"])
	if err != nil {
		log.Fatal(err)
	}
	MongoSession = session
}
