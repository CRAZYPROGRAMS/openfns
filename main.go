package main

import (
	"io/ioutil"
	"log"
	"path"
	"strings"

	"github.com/crazyprograms/openfns/fias"
)

const dataPath string = "./data/"

func main() {
	files, err := ioutil.ReadDir(dataPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.HasPrefix(strings.ToUpper(file.Name()), "AS_ADDROBJ") {
			fias.LoadAddressObjects(path.Join(dataPath, file.Name()))
		}
	}
}
