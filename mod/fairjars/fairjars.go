package fairjars

import (
	"fmt"
	"strings"

	"github.com/Danice123/7mods/mod"
	"github.com/Danice123/7mods/sevenxml"
)

var EXCEPTIONS = map[string]bool{
	"drinkJarBlackStrapCoffee":  true,
	"thrownAmmoMolotovCocktail": true,
}

var TEMPLATE_ITEM = sevenxml.Item{
	Name: "unpackJarTemplate",
	Properties: []*sevenxml.Property{
		{
			Name:  "ItemTypeIcon",
			Value: "treasure",
		},
		{
			Name:  "HoldType",
			Value: "45",
		},
		{
			Name:  "Meshfile",
			Value: "#Other/Items?Misc/sackPrefab.prefab",
		},
		{
			Name:  "DropMeshfile",
			Value: "#Other/Items?Misc/sack_droppedPrefab.prefab",
		},
		{
			Name:  "Material",
			Value: "Morganic",
		},
		{
			Name:  "SellableToTrader",
			Value: "false",
		},
		{
			Name:  "Stacknumber",
			Value: "10",
		},
		{
			Class: "Action0",
			Properties: []*sevenxml.Property{
				{
					Name:  "Class",
					Value: "OpenBundle",
				},
				{
					Name:  "Delay",
					Value: "0",
				},
				{
					Name:  "Use_time",
					Value: "0",
				},
				{
					Name:  "Sound_start",
					Value: "close_garbage",
				},
			},
		},
	},
}

type JarItemEntry struct {
	Tag            string
	Group          string
	UnlockedBy     string
	CustomIconTint string
	StackNumber    string

	CountJars int
	CountItem int
}

func FairJarsMod() *mod.Mod {
	recipes := sevenxml.ReadRecipesXml()
	jarMap := map[string]*JarItemEntry{}
	for _, recipe := range recipes.Recipes {
		if _, ok := EXCEPTIONS[recipe.Name]; ok {
			continue
		}

		jarCount := 0
		for _, ingredients := range recipe.Ingredients {
			if strings.HasPrefix(ingredients.Name, "drinkJar") {
				jarCount++
			}
		}
		if jarCount > 0 {
			jarMap[recipe.Name] = &JarItemEntry{
				CountJars: jarCount,
				CountItem: recipe.Count,
			}
		}
	}

	items := sevenxml.ReadItemsXml()

nextItem:
	for _, item := range items.Items {
		if _, ok := jarMap[item.Name]; ok {
			for _, property := range item.Properties {
				if property.Class == "Action0" {
					for _, subprop := range property.Properties {
						if subprop.Name == "Create_item" && subprop.Value == "drinkJarEmpty" {
							delete(jarMap, item.Name)
							continue nextItem
						}
					}
				}

				if property.Name == "Tags" {
					jarMap[item.Name].Tag = property.Value
				}
				if property.Name == "Group" {
					jarMap[item.Name].Group = property.Value
				}
				if property.Name == "UnlockedBy" {
					jarMap[item.Name].UnlockedBy = property.Value
				}
				if property.Name == "CustomIconTint" {
					jarMap[item.Name].CustomIconTint = property.Value
				}
				if property.Name == "Stacknumber" {
					jarMap[item.Name].StackNumber = property.Value
				}
			}

			if jarMap[item.Name].Tag == "" || jarMap[item.Name].Group == "" {
				if strings.HasPrefix(item.Name, "medical") || strings.HasPrefix(item.Name, "drug") {
					jarMap[item.Name].Tag = "medical"
					jarMap[item.Name].Group = "Science,Medical"
				}
			}
		}
	}

	itemModFile := mod.NewModFile(mod.ITEMS)
	recipeModFile := mod.NewModFile(mod.RECIPES)

	modItemList := []interface{}{}
	modItemList = append(modItemList, &TEMPLATE_ITEM)

	for item, entry := range jarMap {
		modItem := &sevenxml.Item{
			Name: item + "Bundle",
			Properties: []*sevenxml.Property{
				{
					Name:  "Extends",
					Value: TEMPLATE_ITEM.Name,
				},
				{
					Name:  "Tags",
					Value: entry.Tag,
				},
				{
					Name:  "DescriptionKey",
					Value: item + "Desc",
				},
				{
					Name:  "CustomIcon",
					Value: item,
				},
				{
					Name:  "Group",
					Value: entry.Group,
				},
				{
					Class: "Action0",
					Properties: []*sevenxml.Property{
						{
							Name:  "Create_item",
							Value: item + ",drinkJarEmpty",
						},
						{
							Name:  "Create_item_count",
							Value: fmt.Sprintf("%d,%d", entry.CountItem, entry.CountJars),
						},
					},
				},
			},
		}
		if entry.UnlockedBy != "" {
			modItem.Properties = append(modItem.Properties, &sevenxml.Property{
				Name:  "UnlockedBy",
				Value: entry.UnlockedBy,
			})
		}
		if entry.CustomIconTint != "" {
			modItem.Properties = append(modItem.Properties, &sevenxml.Property{
				Name:  "CustomIconTint",
				Value: entry.CustomIconTint,
			})
		}
		if entry.StackNumber != "" {
			modItem.Properties = append(modItem.Properties, &sevenxml.Property{
				Name:  "Stacknumber",
				Value: entry.StackNumber,
			})
		}
		modItemList = append(modItemList, modItem)

		itemModFile.SetAttribute(fmt.Sprintf("/items/item[@name='%sSchematic']/property[@name='Unlocks']", item), "value", item+"Bundle")
		itemModFile.SetAttribute(fmt.Sprintf("/items/item[@name='%sSchematic']/effect_group/triggered_effect[@action='ModifyCVar']", item), "cvar", item+"Bundle")

		recipeModFile.SetAttribute(fmt.Sprintf("/recipes/recipe[@name='%s' and not @tags='salvageScrap']", item), "name", item+"Bundle")
	}

	itemModFile.Append("/items", modItemList)

	return &mod.Mod{
		Name: "Fair_Jars",
		Files: []*mod.Modfile{
			itemModFile,
			recipeModFile,
		},
	}
}
