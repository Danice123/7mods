package main

import (
	"fmt"
	"strconv"
)

func BetterFoodMod() *Modfile {
	items := ReadItemsXml()

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
