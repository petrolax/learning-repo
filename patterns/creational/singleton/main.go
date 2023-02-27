package main

import (
	"fmt"

	"github.com/petrolax/learning-repo/patterns/creational/singleton/properties"
)

type Document struct {
	name     string
	properts properties.Properties
}

func newDocument(name string, properts properties.Properties) Document {
	return Document{
		name: name,
		properts, properts,
	}
}

func main() {
	txtFile := newDocument("file.txt", properties.NewProperties())
	xmlFile := newDocument("file.xml", properties.NewProperties())

	fmt.Println(txtFile.properts == xmlFile.properts)
}
