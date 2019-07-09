package mongo

import (
	"github.com/globalsign/mgo"
)

// MongoStore DB Sessions are maintained inside a struct for better caching of the data stores
// developed based on the stackoverflow answer:
// http://stackoverflow.com/questions/26574594/best-practice-to-maintain-a-mgo-session
type MongoStore struct {
	DBName  string
	session *mgo.Session
}

// Clone the master session and return
func (ms *MongoStore) getSession() *mgo.Session {
	return ms.session.Copy()
}

//GetSessionCollection gets the appropriate MongoDB collection
func (ms *MongoStore) GetSessionCollection(collection string) (*mgo.Session, *mgo.Collection) {
	s := ms.getSession()
	c := s.DB(ms.DBName).C(collection)
	return s, c
}

// Mongo would hold the MongoDB data store
var Mongo MongoStore
