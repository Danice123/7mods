package sevenxml

import "encoding/xml"

type Entities struct {
	XMLName xml.Name `xml:"entity_classes"`
	Items   []*Item  `xml:"item"`
}

type Entity struct {
	XMLName     xml.Name       `xml:"entity_class"`
	Name        string         `xml:"name,attr"`
	Properties  []*Property    `xml:"property"`
	EffectGroup []*EffectGroup `xml:"effect_group"`
}
