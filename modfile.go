package main

import "encoding/xml"

type Modfile struct {
	Filename string
	configs  *Configs
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

func NewModFile(filename string) *Modfile {
	return &Modfile{
		Filename: filename,
		configs: &Configs{
			Sets: []*Set{},
		},
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
