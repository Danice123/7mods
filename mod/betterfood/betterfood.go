package betterfood

import (
	"fmt"
	"strconv"

	"github.com/Danice123/7mods/mod"
	"github.com/Danice123/7mods/sevenxml"
)

func BetterFoodMod() *mod.Mod {
	items := sevenxml.ReadItemsXml()

	itemOverride := mod.NewModFile(mod.ITEMS)
	for _, item := range items.Items {
		for _, eg := range item.EffectGroups {
			for _, effect := range eg.TriggeredEffects {
				if effect.CVar == "$foodAmountAdd" {
					xpath := fmt.Sprintf("/items/item[@name='%s']/effect_group/triggered_effect[@cvar='$foodAmountAdd']/@value", item.Name)
					value, err := strconv.Atoi(effect.Value)
					if err != nil {
						panic(err)
					}
					itemOverride.Set(xpath, strconv.Itoa(value*2))
				}
				if effect.CVar == "$waterAmountAdd" {
					xpath := fmt.Sprintf("/items/item[@name='%s']/effect_group/triggered_effect[@cvar='$waterAmountAdd']/@value", item.Name)
					value, err := strconv.Atoi(effect.Value)
					if err != nil {
						panic(err)
					}
					itemOverride.Set(xpath, strconv.Itoa(value*2))
				}
			}
		}
	}

	return &mod.Mod{
		Name:        "Better_Food",
		Description: "Make food doublegood",
		Files: []*mod.Modfile{
			itemOverride,
		},
	}
}
