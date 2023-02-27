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
		name:     name,
		properts: properts,
	}
}

func main() {
	txtFile := newDocument("file.xml", properties.NewProperties("properties-txt"))
	xmlFile := newDocument("file.txt", properties.NewProperties("properties-xml"))

	fmt.Println(txtFile.properts.Name)
	fmt.Println(xmlFile.properts.Name)
	fmt.Println(txtFile.properts.Name == xmlFile.properts.Name) // true, because properties is singleton
}
