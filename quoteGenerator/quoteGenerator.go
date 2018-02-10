package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kataras/iris"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Quote struct {
	Author  string
	Subject string
	Text    string
}

func main() {

	xmlFile, err := os.Open("Quotes.xml")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully Opened Quotes.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	app := iris.Default()

	session, err := mgo.Dial("localhost")
	if nil != err {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("quote-db").C("quote")
	c.Insert(&Quote{"Lajos Kossuth", "On History", "History is the revelation of providence."})

	app.Get("/", func(ctx iris.Context) {
		result := Quote{}
		err = c.Find(bson.M{"author": "Lajos Kossuth"}).One(&result)
		if err != nil {
			log.Fatal(err)
		}
		ctx.JSON(result)
	})

	app.Run(iris.Addr(":8080"))

}
