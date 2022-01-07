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
	XMLName xml.Name    `xml:"xml"`
	ModInfo *ModInfoTag `xml:"modinfo"`
}

type ModInfoTag struct {
	XMLName     xml.Name      `xml:"modinfo"`
	Name        *TagWithValue `xml:"name"`
	Description *TagWithValue `xml:"description"`
	Author      *TagWithValue `xml:"author"`
	Version     *TagWithValue `xml:"version"`
}

type TagWithValue struct {
	Value string `xml:"value,attr"`
}

// type NameTag struct {
// 	XMLName xml.Name `xml:"name"`
// 	Value   string   `xml:"value,attr"`
// }

// type DescriptionTag struct {
// 	XMLName xml.Name `xml:"description"`
// 	Value   string   `xml:"value,attr"`
// }

// type AuthorTag struct {
// 	XMLName xml.Name `xml:"author"`
// 	Value   string   `xml:"value,attr"`
// }

// type VersionTag struct {
// 	XMLName xml.Name `xml:"version"`
// 	Value   string   `xml:"value,attr"`
// }
