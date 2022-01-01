package main

import (
	"encoding/xml"
	"os"

	"github.com/go-xmlfmt/xmlfmt"
)

func main() {
	mod := BetterFoodMod()

	data, err := xml.Marshal(mod.configs)
	if err != nil {
		panic(err)
	}

	fmtdata := xmlfmt.FormatXML(string(data), "", "\t")
	err = os.WriteFile(mod.Filename, []byte(fmtdata), 0777)
	if err != nil {
		panic(err)
	}
}
