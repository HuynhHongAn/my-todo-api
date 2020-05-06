package database

import "github.com/globalsign/mgo"

type MongoLayer interface {
	Find(condition interface{}) ([]interface{}, error)
}

type mongoLayer struct {
	database string
	collection string
}

func NewMongoLayer (dbName string, collectionName string) MongoLayer {
	return &mongoLayer{
		database: dbName,
		collection: collectionName,
	}
}

func (m mongoLayer) init() (*mgo.Collection, error) {
	url := "mongodb://localhost:27017" //Todo (anhh): move this to env
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	c := session.DB(m.database).C(m.collection)

	return c, nil
}

func (m mongoLayer) Find(condition interface{}) ([]interface{}, error) {
	c, err := m.init()
	if err != nil {
		return nil, err
	}

	var result []interface{}
	if err := c.Find(condition).All(&result); err != nil {
		return nil, err
	}
	return result, nil

}

