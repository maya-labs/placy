package placy

import (
	"io/ioutil"
	"net/http"
)

const (
	_PlacesHost      = "https://maps.googleapis.com/maps/api/place/textsearch/json?"
	_ContryCodesHost = "http://battuta.medunes.net/api/region/"
)

// Config struct
type Init struct {
	GToken      string
	CountryCode string
}

/**
*	This will be the main function for getting the places of all the
*	cities in the contry specified
**/
func GetPlaces(config Init) {
	// TODO
	// Get all the cities of the contry specified
	// Iterate the array of cities
	// Using gplaces for getting the places of the city
	// return the generated array
}

/**
*	IMPLEMENTED BATUTTA API (Open Source)
*	Function that retrieve all the cities using the contrycode
*	of the Init struct
**/
func GetCities(config Init) []byte {
	_api_cities := _ContryCodesHost + config.CountryCode + "/all/?key=c8cd7e1b4cf72844cacceefc1d7f5d0d"

	resp, err := http.Get(_api_cities)

	if err != nil {
		return nil
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	return body
}
