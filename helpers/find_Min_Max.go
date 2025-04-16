package helpers

import (
	"math"
	"sync"

	"groupie/tools"
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
	dd := math.Round((float64(min) + float64(max)) / 2)
	data.Span = int(dd)
	data.MinCrDate = min
	data.MaxCrDate = max
}
