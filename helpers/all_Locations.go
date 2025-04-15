package helpers

import (
	"fmt"
	"sync"

	"groupie/tools"
)

var Locations tools.Index
func AllLocations(allArtists *[]tools.Artists, wg *sync.WaitGroup, data *tools.Data) {
	defer wg.Done()
	err := Fetch("https://groupietrackers.herokuapp.com/api/locations", &Locations)
	if err != nil {
		fmt.Println("Error fetching API:", err)
		return
	}

	locationsSet := make(map[string]bool)

	for _, entry := range Locations.Index {
		for _, loc := range entry.Locations {
			locationsSet[loc] = true
		}
	}
	var locationsList []string
	for location := range locationsSet {
		locationsList = append(locationsList, location)
	}
	data.Locations = locationsList
}
