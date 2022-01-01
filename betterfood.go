package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
)

type Items struct {
	XMLName xml.Name `xml:"items"`
	Items   []*Item  `xml:"item"`
}

type Item struct {
	XMLName     xml.Name     `xml:"item"`
	Name        string       `xml:"name,attr"`
	EffectGroup *EffectGroup `xml:"effect_group"`
}

type EffectGroup struct {
	XMLName          xml.Name           `xml:"effect_group"`
	TriggeredEffects []*TriggeredEffect `xml:"triggered_effect"`
}

type TriggeredEffect struct {
	XMLName   xml.Name `xml:"triggered_effect"`
	Trigger   string   `xml:"trigger,attr"`
	Action    string   `xml:"action,attr"`
	CVar      string   `xml:"cvar,attr"`
	Operation string   `xml:"operation,attr"`
	Value     string   `xml:"value,attr"`
}

func BetterFoodMod() *Modfile {
	data, err := os.ReadFile("items.xml")
	if err != nil {
		panic(err)
	}

	items := Items{}
	err = xml.Unmarshal(data, &items)
	if err != nil {
		panic(err)
	}

	mod := NewModFile("build/Better_Food/Config/items.xml")
	for _, item := range items.Items {
		if item.EffectGroup != nil {
			for _, effect := range item.EffectGroup.TriggeredEffects {
				if effect.CVar == "$foodAmountAdd" {
					xpath := fmt.Sprintf("/items/item[@name='%s']/effect_group/triggered_effect[@cvar='$foodAmountAdd']/@value", item.Name)
					value, err := strconv.Atoi(effect.Value)
					if err != nil {
						panic(err)
					}
					mod.Set(xpath, strconv.Itoa(value*2))
				}
				if effect.CVar == "$waterAmountAdd" {
					xpath := fmt.Sprintf("/items/item[@name='%s']/effect_group/triggered_effect[@cvar='$waterAmountAdd']/@value", item.Name)
					value, err := strconv.Atoi(effect.Value)
					if err != nil {
						panic(err)
					}
					mod.Set(xpath, strconv.Itoa(value*2))
				}
			}
		}
	}

	return mod
}
