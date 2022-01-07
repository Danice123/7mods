package mod

import (
	"fmt"
	"os"
)

type Mod struct {
	Name        string
	Description string
	Files       []Writable
}

type Writable interface {
	Write(string)
}

func (ths *Mod) Write() {
	err := os.MkdirAll(fmt.Sprintf("build/%s/Config", ths.Name), 0777)
	if err != nil {
		panic(err)
	}

	WriteModInfo(ths.Name, ths.Description)

	for _, file := range ths.Files {
		file.Write(ths.Name)
	}
}
