package properties

import (
	"sync"
)

var opt Properties
var once sync.Once

type Properties struct {
	Name           string
	propertiesList []string
}

func NewProperties(name string) Properties {
	once.Do(func() {
		opt = Properties{
			Name:           name,
			propertiesList: []string{"open", "edit", "close"},
		}
	})
	return opt
}
