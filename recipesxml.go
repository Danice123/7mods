package main

import (
	"encoding/xml"
	"os"
)

type Recipes struct {
	XMLName xml.Name  `xml:"recipes"`
	Recipes []*Recipe `xml:"recipe"`
}

type Recipe struct {
	XMLName     xml.Name      `xml:"recipe"`
	Name        string        `xml:"name,attr"`
	Count       int           `xml:"count,attr"`
	Ingredients []*Ingredient `xml:"ingredient"`
}

type Ingredient struct {
	XMLName xml.Name `xml:"ingredient"`
	Name    string   `xml:"name,attr"`
	Count   int      `xml:"count,attr"`
}

func ReadRecipesXml() *Recipes {
	data, err := os.ReadFile("recipes.xml")
	if err != nil {
		panic(err)
	}

	recipes := &Recipes{}
	err = xml.Unmarshal(data, recipes)
	if err != nil {
		panic(err)
	}

	return recipes
}
