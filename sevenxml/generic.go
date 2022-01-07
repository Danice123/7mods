package sevenxml

import "encoding/xml"

type Property struct {
	XMLName    xml.Name    `xml:"property"`
	Name       string      `xml:"name,attr,omitempty"`
	Class      string      `xml:"class,attr,omitempty"`
	Value      string      `xml:"value,attr,omitempty"`
	Properties []*Property `xml:"property"`
}

type EffectGroup struct {
	XMLName          xml.Name  `xml:"effect_group"`
	Name             string    `xml:"name,attr,omitempty"`
	Tiered           string    `xml:"tiered,attr,omitempty"`
	PassiveEffects   []*Effect `xml:"passive_effect"`
	TriggeredEffects []*Effect `xml:"triggered_effect"`
}

type Effect struct {
	Name      string `xml:"name,attr,omitempty"`
	Trigger   string `xml:"trigger,attr,omitempty"`
	Action    string `xml:"action,attr,omitempty"`
	CVar      string `xml:"cvar,attr,omitempty"`
	Operation string `xml:"operation,attr,omitempty"`
	Value     string `xml:"value,attr,omitempty"`
	Tags      string `xml:"tags,attr,omitempty"`
	Exp       string `xml:"exp,attr,omitempty"`
	Level     string `xml:"level,attr,omitempty"`
}
