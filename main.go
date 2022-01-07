package main

import (
	"github.com/Danice123/7mods/mod"
	"github.com/Danice123/7mods/mod/beckerator"
	"github.com/Danice123/7mods/mod/betterfood"
	"github.com/Danice123/7mods/mod/doublegun"
	"github.com/Danice123/7mods/mod/fairjars"
)

var mods = []func() *mod.Mod{
	betterfood.BetterFoodMod,
	doublegun.DoubleDamageGuns,
	fairjars.FairJarsMod,
	beckerator.Beckerator,
}

func main() {
	for _, m := range mods {
		m().Write()
	}
}
