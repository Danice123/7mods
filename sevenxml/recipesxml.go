package sevenxml

import (
	"encoding/xml"
	"os"
)

type RecipeXML struct {
	XMLName xml.Name  `xml:"recipes"`
	Recipes []*Recipe `xml:"recipe"`
}

type Recipe struct {
	XMLName      xml.Name       `xml:"recipe"`
	Name         string         `xml:"name,attr"`
	Count        int            `xml:"count,attr"`
	CraftArea    string         `xml:"craft_area,attr,omitempty"`
	Tags         string         `xml:"tags,attr,omitempty"`
	Ingredients  []*Ingredient  `xml:"ingredient"`
	EffectGroups []*EffectGroup `xml:"effect_group"`
}

type Ingredient struct {
	XMLName xml.Name `xml:"ingredient"`
	Name    string   `xml:"name,attr"`
	Count   int      `xml:"count,attr"`
}

func ReadRecipesXml() *RecipeXML {
	data, err := os.ReadFile("xml/recipes.xml")
	if err != nil {
		panic(err)
	}

	recipes := &RecipeXML{}
	err = xml.Unmarshal(data, recipes)
	if err != nil {
		panic(err)
	}

	return recipes
}
