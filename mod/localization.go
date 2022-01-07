package mod

import (
	"fmt"
	"os"
)

const header = "Key,File,Type,english"

type Localization struct {
	Entries []Entry
}

type Entry struct {
	Id      string
	File    string
	Type    string
	English string
}

func NewLocalizationFile() *Localization {
	return &Localization{
		Entries: []Entry{},
	}
}

func (ths *Localization) Write(modName string) {
	data := header
	for _, entry := range ths.Entries {
		data = fmt.Sprintf("%s\n%s,%s,%s,\"%s\"", data, entry.Id, entry.File, entry.Type, entry.English)
	}

	err := os.WriteFile(fmt.Sprintf("build/%s/Config/Localization.txt", modName), []byte(data), 0777)
	if err != nil {
		panic(err)
	}
}
