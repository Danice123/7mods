package beckerator

import (
	"strconv"

	"github.com/Danice123/7mods/mod"
	"github.com/Danice123/7mods/sevenxml"
)

func buildBeckerator(auger *sevenxml.Item) *sevenxml.Item {
	beckerator := &sevenxml.Item{
		Name:       "meleeToolPickT4Beckerator",
		Properties: []*sevenxml.Property{},
		EffectGroups: []*sevenxml.EffectGroup{
			{
				Name:           "meleeToolPickT4Beckerator",
				PassiveEffects: []*sevenxml.Effect{},
			},
		},
	}

	for _, prop := range auger.Properties {
		if prop.Name == "UnlockedBy" {
			beckerator.Properties = append(beckerator.Properties, &sevenxml.Property{
				Name:  "UnlockedBy",
				Value: "meleeToolPickT4BeckeratorSchematic",
			})
		} else {
			beckerator.Properties = append(beckerator.Properties, prop)
		}
	}

	for _, effect := range auger.EffectGroups[0].PassiveEffects {
		// <passive_effect name="BlockDamage" operation="base_set" value="20.7" tags="perkMiner69r,miningTool"/>
		if effect.Name == "BlockDamage" && effect.Operation == "base_set" {
			v, err := strconv.ParseFloat(effect.Value, 64)
			if err != nil {
				panic(err)
			}

			beckerator.EffectGroups[0].PassiveEffects = append(beckerator.EffectGroups[0].PassiveEffects, &sevenxml.Effect{
				Name:      "BlockDamage",
				Operation: "base_set",
				Value:     strconv.FormatFloat(v*2, 'f', 1, 64),
				Tags:      "perkMiner69r,miningTool",
			})
		} else {
			beckerator.EffectGroups[0].PassiveEffects = append(beckerator.EffectGroups[0].PassiveEffects, effect)
		}
	}

	return beckerator
}

func buildSchematic() *sevenxml.Item {
	return &sevenxml.Item{
		Name: "meleeToolPickT4BeckeratorSchematic",
		Properties: []*sevenxml.Property{
			{
				Name:  "Extends",
				Value: "schematicMaster",
			},
			{
				Name:  "CreativeMode",
				Value: "Player",
			},
			{
				Name:  "CustomIcon",
				Value: "meleeToolPickT3Auger",
			},
			{
				Name:  "Unlocks",
				Value: "meleeToolPickT4Beckerator",
			},
		},
		EffectGroups: []*sevenxml.EffectGroup{
			{
				Tiered: "false",
				TriggeredEffects: []*sevenxml.Effect{
					{
						Trigger:   "onSelfPrimaryActionEnd",
						Action:    "ModifyCVar",
						CVar:      "meleeToolPickT4Beckerator",
						Operation: "set",
						Value:     "1",
					},
					{
						Trigger: "onSelfPrimaryActionEnd",
						Action:  "GiveExp",
						Exp:     "50",
					},
				},
			},
		},
	}
}

func buildRecipe() *sevenxml.Recipe {
	return &sevenxml.Recipe{
		Name:      "meleeToolPickT4Beckerator",
		Count:     1,
		CraftArea: "workbench",
		Tags:      "learnable,perkMiner69r",
		Ingredients: []*sevenxml.Ingredient{
			{
				Name:  "meleeToolAxeT3ChainsawParts",
				Count: 10,
			},
			{
				Name:  "vehicleBicycleHandlebars",
				Count: 1,
			},
			{
				Name:  "smallEngine",
				Count: 2,
			},
			{
				Name:  "resourceForgedSteel",
				Count: 20,
			},
		},
		EffectGroups: []*sevenxml.EffectGroup{
			{
				PassiveEffects: []*sevenxml.Effect{
					{
						Name:      "CraftingIngredientCount",
						Operation: "perc_add",
						Value:     ".5,2.5",
						Tags:      "meleeToolAxeT3ChainsawParts,resourceForgedSteel",
					},
				},
			},
		},
	}
}

func Beckerator() *mod.Mod {
	itemxml := sevenxml.ReadItemsXml()

	var auger *sevenxml.Item
	for _, item := range itemxml.Items {
		if item.Name == "meleeToolPickT3Auger" {
			auger = item
			break
		}
	}

	itemmod := mod.NewModFile(mod.ITEMS)
	itemmod.Append("/items", []interface{}{
		buildBeckerator(auger),
		buildSchematic(),
	})

	recipemod := mod.NewModFile(mod.RECIPES)
	recipemod.Append("/recipes", []interface{}{
		buildRecipe(),
	})

	lootmod := mod.NewModFile(mod.LOOT)
	lootmod.Append("/lootcontainers/lootgroup[@name='schematicsToolsT2']", []interface{}{
		sevenxml.LootGroupItem{
			Name: "meleeToolPickT4BeckeratorSchematic",
		},
	})

	return &mod.Mod{
		Name:        "Beckerator",
		Description: "Gotta mine fast",
		Files: []*mod.Modfile{
			itemmod,
			recipemod,
			lootmod,
		},
	}
}