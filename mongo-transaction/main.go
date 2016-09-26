package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"fmt"
)

type Doc struct {
	Title    string     `bson:"title"`
	Content  string     `bson:"content"`
	Comments []Comments `bson:"comments"`
}

type Comments struct {
	Name    string `bson:"Name"`
	Content string `bson:"content"`
}

func main() {
	session, err := mgo.Dial("192.168.99.100:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("people")
	err = c.Insert(
		&Doc{
			Title:   "title1",
			Content: "content1",
			Comments: []Comments{
				Comments{Name: "title1-name1", Content: "title1-email1"},
				Comments{Name: "title1-name2", Content: "title1-email2"},
			},
		},
		&Doc{
			Title:   "title2",
			Content: "content2",
			Comments: []Comments{
				Comments{Name: "title2-name1", Content: "title2-email1"},
				Comments{Name: "title2-name2", Content: "title2-email2"},
			},
		})
	if err != nil {
		panic(err)
	}
	result := Doc{}
	err = c.Find(bson.M{"title": "title1"}).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println("result:", result)

	// 再查一次 条件为 {"title": "title2", "comments": {"$elemMatch": {"content": "title2-email2"}}}  这个条件在用mongo 是能查到记录的
	// 这个运行时会报 missing type in composite literal
	result2 := Doc{}
	err = c.Find(bson.M{"title": "title2", "comments": bson.M{"$elemMatch": bson.M{"content": "title2-email2"}}}).One(&result2)
	if err != nil {
		panic(err)
	}
	fmt.Println("result2:", result2)

}
