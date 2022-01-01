package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Items struct {
	XMLName xml.Name `xml:"items"`
	Items   []*Item  `xml:"item"`
}

type Item struct {
	XMLName xml.Name `xml:"item"`
	Name    string   `xml:"name,attr"`
}

func BetterFoodMod() {
	data, err := os.ReadFile("items.xml")
	if err != nil {
		panic(err)
	}

	items := []Items{}
	err = xml.Unmarshal(data, &items)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", items)
}
