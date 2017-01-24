package models

import (
	"github.com/johnsudaar/moviie_api/configuration"
	"gopkg.in/errgo.v1"
	"gopkg.in/mgo.v2/bson"
)

type Share struct {
	From    string `json:"from" bson:"from"`
	To      string `json:"to" bson:"to"`
	Movie   Movie  `json:"movie" bson:"movie"`
	MovieID int64  `json:"-" bson:"movie_id"`
}

func (s *Share) Save() error {
	c := configuration.MongoSession.DB(configuration.E["MONGO_DB"]).C("user")

	var res *Share = nil

	err := c.Find(bson.M{
		"from":     s.From,
		"to":       s.To,
		"movie_id": s.MovieID,
	}).One(res)

	if err != nil && err.Error() != "not found" {
		return errgo.Mask(err)
	}

	if res != nil {
		return errgo.New("User already present")
	}

	err = c.Insert(s)

	if err != nil {
		return errgo.Mask(err)
	}
	return nil
}

func FindShareByUser(user string) ([]Share, error) {
	c := configuration.MongoSession.DB(configuration.E["MONGO_DB"]).C("user")

	var res []Share
	res = make([]Share, 0)
	err := c.Find(bson.M{
		"to": user,
	}).All(&res)

	if err != nil {
		return nil, errgo.Mask(err)
	}

	return res, nil
}
