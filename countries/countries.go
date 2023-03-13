package countries

import (
	"errors"
	"fmt"
)

var myCountry string = "Japan"
var collection []string
var errorNotFound error = errors.New("country not found")

// Add Agrega el valor de un pais a la coleccion de paises
func Add(country string) {
	collection = append(collection, country)
}

// SetMyCountry Selecciona un pais como pais favorito.
func SetMyCountry(country string) error {
	if !isIncollection(country) {
		return errorNotFound
	}

	myCountry = country
	return nil
}

// List Lista todos los paices
func List() {
	for i, c := range collection {
		myCountryMsg := ""
		if c == myCountry {
			myCountryMsg = "[My Country]"
		}
		fmt.Printf("%d. %s %s\n", i+1, c, myCountryMsg)
	}
}

func isIncollection(country string) bool {
	for _, c := range collection {
		if c == country {
			return true
		}
	}

	return false
}
