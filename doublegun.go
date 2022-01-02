package main

import (
	"fmt"
	"strconv"
	"strings"
)

func DoubleDamageGuns() []*Modfile {
	items := ReadItemsXml()

	mod := NewModFile("build/Double_Gun_Damage/Config/items.xml")

	for _, item := range items.Items {
		if strings.HasPrefix(item.Name, "ammo") && !strings.Contains(item.Name, "Bundle") && !strings.HasPrefix(item.Name, "ammoProjectile") && !strings.HasPrefix(item.Name, "ammoGas") {

			for _, eg := range item.EffectGroup {
				if eg.Name == item.Name {
					for _, pe := range eg.PassiveEffects {
						if pe.Name == "EntityDamage" && pe.Operation == "base_set" {
							damage, err := strconv.ParseFloat(pe.Value, 64)
							if err != nil {
								panic(err)
							}
							mod.SetAttribute(fmt.Sprintf("/items/item[@name='%s']/effect_group[@name='%s']/passive_effect[@name='EntityDamage' and @operation='base_set']", item.Name, item.Name), "value", strconv.Itoa(int(damage*2)))
						}
					}
				}
			}
		}
	}

	return []*Modfile{mod}
}
