package config

import (
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"log"
	"os"
)

var mongodbSession *mgo.Session
var (
	dbPrefix = "memorize"
)

// DBSession returns the current db session.
func DBSession() *mgo.Session {
	if mongodbSession != nil {
		return mongodbSession
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		uri = "mongodb://localhost"
	}

	var err error
	mongodbSession, err = mgo.Dial(uri)
	if mongodbSession == nil || err != nil {
		log.Fatalf("Can't connect to mongo, go error %v\n", err)
	}

	mongodbSession.SetSafe(&mgo.Safe{})
	return mongodbSession
}

// DB returns a database given a name.
func DB(name string) *mgo.Database {
	return DBSession().DB(name)
}

// DefaultDB returns the default database.
func DefaultDB() *mgo.Database {
	switch Environment {
	case "test":
		{
			return DB(dbPrefix + "-test")
		}
	case "production":
		{
			return DB(dbPrefix + "-production")
		}
	}

	return DB(dbPrefix + "-development")
}

// AddBasicIndex add a ascending index given a list of `keys`. The index is always built in background.
func AddBasicIndex(collection *mgo.Collection, keys ...string) {
	collection.EnsureIndex(mgo.Index{
		Key:        keys,
		Background: true,
	})
}

// vi:syntax=go
