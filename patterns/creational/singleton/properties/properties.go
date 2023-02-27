package properties

import "sync"

var opt Properties
var once sync.Once

type Properties struct {
	propertiesList []string
}

func NewProperties() Properties {
	once.Do(func() {
		opt = Properties{
			propertiesList: []string{"open", "edit", "close"},
		}
	})
	return opt
}
