package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type GameData struct {
	BaseGameText BaseGameText
}

type BaseGameText struct {
	XMLName xml.Name
	Rows    []Row `xml:"Row"`
}

type Row struct {
	Subject string `xml:"Tag,attr"`
	Text    string
}

func main() {

	xmlFile, err := os.Open("quote.xml")

	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	byteValue, err := ioutil.ReadAll(xmlFile)

	if err != nil {
		log.Fatal(err)
	}

	var gameData GameData
	xml.Unmarshal(byteValue, &gameData)

	for i := 0; i < len(gameData.BaseGameText.Rows); i++ {
		fmt.Println("Quote Data: " + gameData.BaseGameText.Rows[i].Text)
	}

}
