package handler

import (
	"groupie/helpers"
	"groupie/tools"
	"sync"
)

func Handle_data(allArtists *[]tools.Artists, data *tools.Data) {
	var wg sync.WaitGroup
	wg.Add(2)
	go helpers.MinMax(allArtists, &wg, data)

	go helpers.AllLocations(allArtists, &wg, data)
	wg.Wait()
}
