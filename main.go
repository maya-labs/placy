package main

import "placy/placy"

func main() {
	println("hello")

	config := placy.Init{
		GToken:      "",
		CountryCode: "mx",
	}

	placy.GetCities(config)

}
