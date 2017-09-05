package xmlloader

import (
	"encoding/xml"
	"log"
	"os"
	"reflect"
)

func LoadStream(filename string, tagName string, defaultItem interface{}, streamitem func(item interface{}) error) error {
	itemTypeElem := reflect.TypeOf(defaultItem).Elem()
	itemValue := reflect.ValueOf(defaultItem).Elem()
	f, err := os.OpenFile(filename, os.O_RDONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	decoder := xml.NewDecoder(f)
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		// Inspect the type of the token just read.
		switch se := t.(type) {
		case xml.StartElement:
			// If we just read a StartElement token
			if se.Name.Local == tagName {
				vValue := reflect.New(itemTypeElem)
				vValue.Elem().Set(itemValue)
				v := vValue.Interface()

				if err = decoder.DecodeElement(v, &se); err != nil {
					return err
				}
				// Do some stuff with the page.
				if err = streamitem(v); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
