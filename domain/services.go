package domain

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getCollection() *mgo.Collection {
	host := []string{
		"localhost",
	}
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs: host,
	})
	if err != nil {
		panic(err)
	}
	collection := session.DB("example").C("todos")
	return collection

}

func all() []Todo {
	collection := getCollection()
	var todos []Todo
	collection.Find(nil).All(&todos)

	return todos
}
func saveTodo(todo Todo) {

	collection := getCollection()

	if err := collection.Insert(todo); err != nil {
		panic(err)
	}
}

func findById(id string) Todo {

	var todo Todo
	collection := getCollection()
	err := collection.Find(bson.M{"id": id}).One(&todo)
	if err != nil {
		panic(err)
	}
	return todo
}
