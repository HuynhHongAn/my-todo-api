package database

import (
	"github.com/globalsign/mgo"
	"my-todo-api/schema"
)

type MongoLayer interface {
	Find(filterParam schema.FilterParam) ([]interface{}, error)
	Insert(data interface{}) error
	Update(condition interface{}, data interface{}) error
	UpdateAll(condition interface{}, data interface{}) error
	Remove(condition interface{}) error
	RemoveAll(condition interface{}) error
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

func (m mongoLayer) Find(filterParam schema.FilterParam) ([]interface{}, error) {
	c, err := m.init()
	if err != nil {
		return nil, err
	}

	var result []interface{}
	if err := c.Find(filterParam.Condition).Skip(filterParam.Skip).Limit(filterParam.Limit).Sort(filterParam.Sort).All(&result); err != nil {
		return nil, err
	}
	return result, nil

}

func (m mongoLayer) Insert(data interface{}) error {
	c, err := m.init()
	if err != nil {
		return err
	}

	return c.Insert(data)
}

func (m mongoLayer) Update(condition interface{}, data interface{}) error {
	c, err := m.init()
	if err != nil {
		return err
	}

	return c.Update(condition, data)
}

func (m mongoLayer) UpdateAll(condition interface{}, data interface{}) error {
	c, err := m.init()
	if err != nil {
		return err
	}

	_, err = c.UpdateAll(condition, data)
	return err
}

func (m mongoLayer) Remove(condition interface{}) error {
	c, err := m.init()
	if err != nil {
		return err
	}

	return c.Remove(condition)
}

func (m mongoLayer) RemoveAll(condition interface{}) error {
	c, err := m.init()
	if err != nil {
		return err
	}

	_, err = c.RemoveAll(condition)
	return err
}