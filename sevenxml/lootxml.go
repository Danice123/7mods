package sevenxml

import (
	"encoding/xml"
	"os"
)

type LootXML struct {
	XMLName    xml.Name `xml:"lootcontainers"`
	LootGroups []*LootGroup
}

type LootGroup struct {
	XMLName xml.Name `xml:"lootgroup"`
	Name    string   `xml:"name,attr"`
	Items   []*LootGroupItem
}

type LootGroupItem struct {
	XMLName xml.Name `xml:"item"`
	Name    string   `xml:"name,attr"`
}

func ReadLootXml() *LootXML {
	data, err := os.ReadFile("xml/loot.xml")
	if err != nil {
		panic(err)
	}

	loot := &LootXML{}
	err = xml.Unmarshal(data, loot)
	if err != nil {
		panic(err)
	}

	return loot
}
