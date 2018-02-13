package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode"
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

type Quote struct {
	Subject string
	Text    string
	Author  string
}

func makeQuote(row Row) Quote {

	x := func(c rune) bool {
		return !unicode.IsLetter(c)
	}

	strArray := strings.FieldsFunc(row.Subject, x)
	strArray = strArray[2:]
	strArray = strArray[:len(strArray)-1]
	subject := strings.Join(strArray, " ")
	subject = strings.Title(strings.ToLower(subject))
	strArray = strings.Split(row.Text, "[NEWLINE]– ")
	author := strArray[1]
	text := strArray[0]
	text = strings.Replace(text, "”", "", -1)
	text = strings.Replace(text, "“", "", -1)
	quote := Quote{
		Subject: subject,
		Author:  author,
		Text:    text}
	return quote
}

func main() {

	xmlFile, err := os.Open("quotes.xml")

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
		quote := makeQuote(gameData.BaseGameText.Rows[i])
		b, err := json.Marshal(quote)
		if err != nil {
			fmt.Println("error:", err)
		}
		os.Stdout.Write(b)
	}

}
