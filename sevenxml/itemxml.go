package sevenxml

import (
	"encoding/xml"
	"os"
)

type Items struct {
	XMLName xml.Name `xml:"items"`
	Items   []*Item  `xml:"item"`
}

type Item struct {
	XMLName     xml.Name       `xml:"item"`
	Name        string         `xml:"name,attr"`
	Properties  []*Property    `xml:"property"`
	EffectGroup []*EffectGroup `xml:"effect_group"`
}

type Property struct {
	XMLName    xml.Name    `xml:"property"`
	Name       string      `xml:"name,attr,omitempty"`
	Class      string      `xml:"class,attr,omitempty"`
	Value      string      `xml:"value,attr,omitempty"`
	Properties []*Property `xml:"property"`
}

type EffectGroup struct {
	XMLName          xml.Name           `xml:"effect_group"`
	Name             string             `xml:"name,attr,omitempty"`
	PassiveEffects   []*PassiveEffect   `xml:"passive_effect"`
	TriggeredEffects []*TriggeredEffect `xml:"triggered_effect"`
}

type PassiveEffect struct {
	XMLName   xml.Name `xml:"passive_effect"`
	Name      string   `xml:"name,attr"`
	Operation string   `xml:"operation,attr"`
	Value     string   `xml:"value,attr"`
}

type TriggeredEffect struct {
	XMLName   xml.Name `xml:"triggered_effect"`
	Trigger   string   `xml:"trigger,attr"`
	Action    string   `xml:"action,attr"`
	CVar      string   `xml:"cvar,attr"`
	Operation string   `xml:"operation,attr"`
	Value     string   `xml:"value,attr"`
}

func ReadItemsXml() *Items {
	data, err := os.ReadFile("items.xml")
	if err != nil {
		panic(err)
	}

	items := &Items{}
	err = xml.Unmarshal(data, items)
	if err != nil {
		panic(err)
	}

	return items
}
