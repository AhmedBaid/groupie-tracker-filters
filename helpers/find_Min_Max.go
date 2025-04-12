package helpers

import (
	"groupie/tools"
	"sync"
)

func MinMax(allArtists *[]tools.Artists, wg *sync.WaitGroup, data *tools.Data) {
	defer wg.Done()
	// found maxCrDate and minCrDate
	var slice []int
	for _, artist := range *allArtists {
		slice = append(slice, artist.CreationDate)
	}
	min, max := slice[0], slice[0]
	for _, date := range slice[1:] {
		if date < min {
			min = date
		}
		if date > max {
			max = date
		}
	}
	data.MinCrDate = min
	data.MaxCrDate = max
}
