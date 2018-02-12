package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Row struct {
	Subject string `xml:"tag,attr"`
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

	var row Row
	xml.Unmarshal(byteValue, &row)

	fmt.Println(row)

}
