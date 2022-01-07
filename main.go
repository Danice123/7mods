package main

import (
	"github.com/Danice123/7mods/mod"
	"github.com/Danice123/7mods/mod/betterfood"
	"github.com/Danice123/7mods/mod/doublegun"
)

var mods = []func() *mod.Mod{
	betterfood.BetterFoodMod,
	doublegun.DoubleDamageGuns,
}

func main() {
	for _, m := range mods {
		m().Write()
	}
}
