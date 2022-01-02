package main

import (
	"encoding/xml"
	"os"

	"github.com/go-xmlfmt/xmlfmt"
)

func main() {
	mod := DoubleDamageGuns()

	for _, file := range mod {
		data, err := xml.Marshal(file.configs)
		if err != nil {
			panic(err)
		}

		fmtdata := xmlfmt.FormatXML(string(data), "", "\t")
		err = os.WriteFile(file.Filename, []byte(fmtdata), 0777)
		if err != nil {
			panic(err)
		}
	}

}
