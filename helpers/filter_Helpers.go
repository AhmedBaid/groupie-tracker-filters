package helpers

import (
	"groupie/tools"
	"strconv"
	"strings"
)

func ExtractYear(album string) int {
	if len(album) < 4 {
		return 0
	}
	year, err := strconv.Atoi(album[:4])
	if err != nil {
		return 0
	}
	return year
}

func MemberMatch(n int, members []string) bool {
	for _, m := range members {
		val, err := strconv.Atoi(m)
		if err == nil && val == n {
			return true
		}
	}
	return false
}

func HasLocation(artist tools.Artists, loc string) bool {
	var locationData *tools.LocationDataFilter
	err := Fetch_By_Id(artist.Locations, &locationData)
	if err != nil {
		return false
	}
	for _, l := range locationData.Locations {
		if strings.Contains(strings.ToLower(l), strings.ToLower(loc)) {
			return true
		}
	}
	return false
}
