package main

import (
	"fmt"
	"github.com/carloshjoaquim/golang-testing/src/api/providers/locations_provider"
)

func main() {
	country, err := locations_provider.GetCountry("BR")

	fmt.Println(err)
	fmt.Println(country)
}