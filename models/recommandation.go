package models

import (
	"github.com/johnsudaar/moviie_api/configuration"
	"gopkg.in/errgo.v1"
	"gopkg.in/mgo.v2/bson"
)

type Recommendation struct {
	From string `bson:"from" json:"from"`
	To   string `bson:"from" json:"to"`
	Id   string `bson:"film_id" json:"id"`
}

func (r *Recommendation) Save() error {
	c := configuration.MongoSession.DB(configuration.E["MONGO_DB"]).C("recommendation")
	n, err := c.Find(bson.M{"from": r.From, "to": r.To, "film_id": r.Id}).Count()

	if err != nil {
		return errgo.Mask(err)
	}

	if n > 0 {
		return errgo.New("You've already recommended this film to this person.")
	}

	err = c.Insert(r)
	if err != nil {
		return errgo.Mask(err)
	}

	return nil
}
