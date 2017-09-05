package fias

import (
	"fmt"

	"github.com/crazyprograms/openfns/xmlloader"
)

type FIASObject struct {
	AOID       string `xml:"AOID,attr"`
	AOGUID     string `xml:"AOGUID,attr"`
	PARENTGUID string `xml:"PARENTGUID,attr"`
	FORMALNAME string `xml:"FORMALNAME,attr"`
	OFFNAME    string `xml:"OFFNAME,attr"`
}

func LoadAddressObjects(filename string) {
	xmlloader.LoadStream(filename, "Object", &FIASObject{}, func(item interface{}) error {
		o := item.(*FIASObject)
		fmt.Println(o)
		return nil
	})
	fmt.Println(filename)
}
