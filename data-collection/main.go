package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
)

const (
	VERSION = "0.0.01.0609"
)

func checkErr(label string, err error) {
	if err != nil {
		fmt.Println(label+":", err)
		os.Exit(1)
	}
}

type Person struct {
	Id_     bson.ObjectId `bson:"_id"`
	Name    string        `bson:"name"`
	Phone   string        `bson:"phone"`
	Account int           `bson:"account"`
}

func main() {
	session, err := mgo.Dial("192.168.99.100")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(&Person{bson.NewObjectId(), "Ale", "+55 53 8116 9639", 300},
		&Person{bson.NewObjectId(), "Cla", "+55 53 8402 8510", 300})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	result1 := Person{}
	err = c.Find(bson.M{"name": "Cla"}).One(&result1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result)

	fmt.Println("Phone:", result.Id_)
	fmt.Println("Phone:", result1.Id_)

	//---------------------------------------------------------------------------------------------------------------------------------//

	runner := txn.NewRunner(c)
	ops := []txn.Op{{
		C:      "people",
		Id:     result.Id_,
		Assert: bson.M{"account": bson.M{"$gte": 100}},
		Update: bson.M{"$inc": bson.M{"account": -100}},
	}, {
		C:      "people",
		Id:     result1.Id_,
		Assert: bson.M{"account": bson.M{"$gte": 100}},
		Update: bson.M{"$inc": bson.M{"account": 100}},
	}}
	id := bson.NewObjectId() // Optional
	err = runner.Run(ops, id, nil)

	fmt.Println("err", err)
}
