package mod

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/go-xmlfmt/xmlfmt"
)

type TargetFile string

const ITEMS = TargetFile("items.xml")
const ENTITYCLASSES = TargetFile("entityclasses.xml")
const PROGRESSION = TargetFile("progression.xml")
const RECIPES = TargetFile("recipes.xml")
const LOOT = TargetFile("loot.xml")

type Modfile struct {
	target  TargetFile
	configs *Configs
}

type Configs struct {
	XMLName       xml.Name        `xml:"configs"`
	Sets          []*Set          `xml:"set"`
	SetAttributes []*SetAttribute `xml:"setattribute"`
	Appends       []*Append       `xml:"append"`
}

type Set struct {
	XMLName xml.Name `xml:"set"`
	XPath   string   `xml:"xpath,attr"`
	Value   string   `xml:",chardata"`
}

type SetAttribute struct {
	XMLName xml.Name `xml:"setattribute"`
	XPath   string   `xml:"xpath,attr"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:",chardata"`
}

type Append struct {
	XMLName xml.Name `xml:"append"`
	XPath   string   `xml:"xpath,attr"`
	Values  []interface{}
}

func NewModFile(target TargetFile) *Modfile {
	return &Modfile{
		target: target,
		configs: &Configs{
			Sets: []*Set{},
		},
	}
}

func (ths *Modfile) Write(modName string) {
	data, err := xml.Marshal(ths.configs)
	if err != nil {
		panic(err)
	}

	fmtdata := xmlfmt.FormatXML(fmt.Sprintf("%s%s", `<?xml version="1.0" encoding="UTF-8"?>`, string(data)), "", "\t")
	err = os.WriteFile(fmt.Sprintf("build/%s/Config/%s", modName, ths.target), []byte(fmtdata)[2:], 0777)
	if err != nil {
		panic(err)
	}
}

func (ths *Modfile) Set(path string, value string) {
	ths.configs.Sets = append(ths.configs.Sets, &Set{
		XPath: path,
		Value: value,
	})
}

func (ths *Modfile) SetAttribute(path string, name string, value string) {
	ths.configs.SetAttributes = append(ths.configs.SetAttributes, &SetAttribute{
		XPath: path,
		Name:  name,
		Value: value,
	})
}

func (ths *Modfile) Append(path string, values []interface{}) {
	ths.configs.Appends = append(ths.configs.Appends, &Append{
		XPath:  path,
		Values: values,
	})
}
