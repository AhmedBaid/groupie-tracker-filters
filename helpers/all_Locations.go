package helpers

import (
	"groupie/tools"
	"strings"
	"sync"
)

func AllLocations(allArtists *[]tools.Artists, wg *sync.WaitGroup, data *tools.Data) {
	// Collect unique locations mn kol artist
	locationsSet := make(map[string]bool)
	var locationsList []string

	for _, artist := range *allArtists {
		var locData *tools.LocationDataFilter
		defer wg.Done()
		err := Fetch_By_Id(artist.Locations, &locData)
		if err != nil {
			continue
		}

		for _, loc := range locData.Locations {
			cleanedLoc := strings.TrimSpace(loc)
			locationsSet[cleanedLoc] = true
		}
	}

	for loc := range locationsSet {
		locationsList = append(locationsList, loc)
	}
	data.Locations = locationsList
}
