package main

import (
	// "encoding/json"
	// "encoding/xml"
	"encoding/xml"
	"fmt"
	"log"
	"os"

	"github.com/kataras/iris"
)

type Quote struct {
	Author  string
	Subject string
	Text    string
}

type GameData struct {
}

type BaseGameText struct {
	XMLName xml.Name `xml:"baseGameText" json:"-"`
	RowList []Row    `xml:"row" json:"Row"`
}

type Row struct {
	Subject string `xml:"Tag,attr"`
	Text    string `xml:"Text"`
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

	fileStat, err := xmlFile.Stat()
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, fileStat.Size())

	part, err := xmlFile.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	// session, err := mgo.Dial("localhost")
	// if nil != err {
	// 	panic(err)
	// }
	// defer session.Close()
	// session.SetMode(mgo.Monotonic, true)

	// c := session.DB("quote-db").C("quote")
	// c.Insert(&Quote{"Lajos Kossuth", "On History", "History is the revelation of providence."})

	app.Get("/", func(ctx iris.Context) {
		// result := Quote{}
		// err = c.Find(bson.M{"author": "Lajos Kossuth"}).One(&result)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// ctx.JSON(result)
		ctx.HTML(string(part))
	})

	app.Run(iris.Addr(":8080"))

}
