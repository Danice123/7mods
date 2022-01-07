package mod

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"

	"github.com/go-xmlfmt/xmlfmt"
)

func WriteModInfo(slug string, description string) {
	file := XMLTag{
		ModInfo: &ModInfoTag{
			Name: &TagWithValue{
				Value: strings.ReplaceAll(slug, "_", " "),
			},
			Description: &TagWithValue{
				Value: description,
			},
			Author: &TagWithValue{
				Value: "danice123",
			},
			Version: &TagWithValue{
				Value: "1.0.0",
			},
		},
	}

	data, err := xml.Marshal(file)
	if err != nil {
		panic(err)
	}

	fmtdata := xmlfmt.FormatXML(fmt.Sprintf("%s%s", `<?xml version="1.0" encoding="UTF-8"?>`, string(data)), "", "\t")
	err = os.WriteFile(fmt.Sprintf("build/%s/ModInfo.xml", slug), []byte(fmtdata)[2:], 0777)
	if err != nil {
		panic(err)
	}
}

type XMLTag struct {
	XMLName xml.Name `xml:"xml"`
	ModInfo *ModInfoTag
}

type ModInfoTag struct {
	XMLName     xml.Name      `xml:"ModInfo"`
	Name        *TagWithValue `xml:"Name"`
	Description *TagWithValue `xml:"Description"`
	Author      *TagWithValue `xml:"Author"`
	Version     *TagWithValue `xml:"Vmersion"`
}

type TagWithValue struct {
	Value string `xml:"value,attr"`
}
