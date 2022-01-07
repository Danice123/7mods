package sevenxml

import (
	"encoding/xml"
	"os"
)

type ItemXML struct {
	XMLName xml.Name `xml:"items"`
	Items   []*Item  `xml:"item"`
}

type Item struct {
	XMLName      xml.Name       `xml:"item"`
	Name         string         `xml:"name,attr"`
	Properties   []*Property    `xml:"property"`
	EffectGroups []*EffectGroup `xml:"effect_group"`
}

func ReadItemsXml() *ItemXML {
	data, err := os.ReadFile("xml/items.xml")
	if err != nil {
		panic(err)
	}

	items := &ItemXML{}
	err = xml.Unmarshal(data, items)
	if err != nil {
		panic(err)
	}

	return items
}
