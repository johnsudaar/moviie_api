package models

import (
	"github.com/johnsudaar/moviie_api/configuration"
	"github.com/satori/go.uuid"
	"gopkg.in/errgo.v1"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Username string   `bson:"username" json:"username"`
	Password string   `bson:"password" json:"-"`
	ApiKey   string   `bson:"api_key" json:"api_key"`
	Friends  []string `bson:"friends" json:"friends"`
	Movies   []*Movie `bson:"movies" json:"movies"`
}

func (u *User) Update() error {
	c := configuration.MongoSession.DB(configuration.E["MONGO_DB"]).C("user")

	err := c.Update(bson.M{
		"username": u.Username,
	}, u)

	if err != nil {
		return errgo.Mask(err)
	}
	return nil
}

func (u *User) Save() error {
	colection := configuration.MongoSession.DB(configuration.E["MONGO_DB"]).C("user")
	var res *User = nil
	err := colection.Find(bson.M{
		"username": u.Username,
	}).One(res)

	if err != nil && err.Error() != "not found" {
		return errgo.Mask(err)
	}

	if res != nil {
		return errgo.New("User already present")
	}

	u.ApiKey = uuid.NewV4().String()

	err = colection.Insert(u)

	if err != nil {
		return errgo.Mask(err)
	}
	return nil
}

func (u *User) AddFriend(friend *User) error {
	err1 := u._addFriend(friend)
	err2 := friend._addFriend(u)

	if err1 != nil {
		return errgo.Mask(err1)
	}

	if err2 != nil {
		return errgo.Mask(err2)
	}
	return nil

}

func (u *User) _addFriend(friend *User) error {
	u.Friends = append(u.Friends, friend.Username)
	c := configuration.MongoSession.DB(configuration.E["MONGO_DB"]).C("user")

	err := c.Update(bson.M{
		"username": u.Username,
	}, u)

	if err != nil {
		return errgo.Mask(err)
	}
	return nil
}

func Login(username, password string) (*User, error) {
	colection := configuration.MongoSession.DB(configuration.E["MONGO_DB"]).C("user")

	res := User{}
	err := colection.Find(bson.M{
		"username": username,
		"password": password,
	}).One(&res)
	if err != nil {
		return nil, errgo.Mask(err)
	}

	return &res, nil
}

func FindUserByUsername(username string) (*User, error) {
	c := configuration.MongoSession.DB(configuration.E["MONGO_DB"]).C("user")
	res := User{}

	err := c.Find(bson.M{
		"username": username,
	}).One(&res)

	if err != nil {
		return nil, errgo.Mask(err)
	}
	return &res, nil
}

func FindUserByApiKey(apiKey string) (*User, error) {
	c := configuration.MongoSession.DB(configuration.E["MONGO_DB"]).C("user")
	res := User{}
	err := c.Find(bson.M{
		"api_key": apiKey,
	}).One(&res)
	if err != nil {
		return nil, errgo.Mask(err)
	}

	return &res, nil
}

func SearchUser(pattern string) ([]string, error) {
	c := configuration.MongoSession.DB(configuration.E["MONGO_DB"]).C("user")
	var result []User
	err := c.Find(bson.M{
		"username": &bson.RegEx{
			Pattern: ".*" + pattern + ".*",
			Options: "i",
		},
	}).All(&result)
	if err != nil {
		return nil, errgo.Mask(err)
	}

	userStrings := make([]string, len(result))
	for i, u := range result {
		userStrings[i] = u.Username
	}
	return userStrings, nil
}
