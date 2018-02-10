package main

import (
	"log"

	"github.com/kataras/iris"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Quote struct {
	Author  string
	Subject string
	Text    string
}

func main() {
	app := iris.Default()

	session, err := mgo.Dial("localhost")
	if nil != err {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("quote-db").C("quote")
	c.Insert(&Quote{"Lajos Kossuth", "On History", "History is the revelation of providence."})

	app.Get("/quote", func(ctx iris.Context) {
		result := Quote{}
		err = c.Find(bson.M{"Author": "Lajos Kossuth"}).One(&result)
		if err != nil {
			log.Fatal(err)
		}
		ctx.JSON(result)
	})

	app.Run(iris.Addr(":8080"))
}
