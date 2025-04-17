package helpers

import (
	"strconv"
	"strings"

	"groupie/tools"
)

func CheckCreationDate(a *tools.Artists, min string, max string) bool {
	minV, _ := strconv.Atoi(min)
	maxV, _ := strconv.Atoi(max)
	if (minV == 1987 && maxV == 1987) || (a.CreationDate >= minV && maxV >= a.CreationDate) {
		return true
	}
	return false
}

func CheckFirstAlbum(a *tools.Artists, year1, year2 string) bool {
	if len(year1) == 0 && len(year2) == 0 {
		return true
	}

	minyear, _ := strconv.Atoi(year1)
	maxyear, _ := strconv.Atoi(year2)
	firstAlbYear := strings.Split(a.FirstAlbum, "-")[2]
	for i := minyear; i <= maxyear; i++ {
		if firstAlbYear == strconv.Itoa(i) {
			return true
		}
	}
	return false
}

func CheckNumberOfMembers(a *tools.Artists, members []string) bool {
	if len(members) == 0 {
		return true
	}

	for _, e := range members {
		nb, _ := strconv.Atoi(e)
		if len(a.Members) == nb {
			return true
		}
	}
	return false
}

func CheckLocations(locations *tools.Index, artists *tools.Artists, location string) bool {
	if location == "" {
		return true
	}
	// if location == "seattle-usa" {
	// 	location = "washington-usa"
	// }
	for _, locations := range locations.Index {
		for _, loc := range locations.Locations {
			if loc == location {
				if locations.ID == artists.Id {
					return true
				}
			}
		}
	}
	return false
}

func ArtistsFiltred(allArtists *[]tools.Artists, minCrStr, maxCrStr, firstAlbumMin, firstAlbumMax, location string, members []string, locations tools.Index) *[]tools.Artists {
	filteredArtists := []tools.Artists{}
	for _, artist := range *allArtists {
		hasDate := CheckCreationDate(&artist, minCrStr, maxCrStr)
		hasFirstAlbum := CheckFirstAlbum(&artist, firstAlbumMin, firstAlbumMax)
		hasMembers := CheckNumberOfMembers(&artist, members)
		hasLocations := CheckLocations(&locations, &artist, location)
		if hasMembers && hasFirstAlbum && hasDate && hasLocations {
			filteredArtists = append(filteredArtists, artist)
		}
	}
	return &filteredArtists
}
