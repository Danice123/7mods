package main

import "encoding/xml"

type Modfile struct {
	Filename string
	configs  *Configs
}

type Configs struct {
	XMLName xml.Name `xml:"configs"`
	Sets    []*Set   `xml:"set"`
}

type Set struct {
	XMLName xml.Name `xml:"set"`
	XPath   string   `xml:"xpath,attr"`
	Value   string   `xml:",chardata"`
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
