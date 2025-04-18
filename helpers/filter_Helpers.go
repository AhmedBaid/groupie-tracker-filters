package helpers

import (
	"strconv"
	"strings"

	"groupie/tools"
)

func CheckCreationDate(allArtists *tools.Artists, minCr string, maxCr string) bool {
	minCrDate, _ := strconv.Atoi(minCr)
	maxCrDate, _ := strconv.Atoi(maxCr)
	if (minCrDate == 1987 && maxCrDate == 1987) || (allArtists.CreationDate >= minCrDate && maxCrDate >= allArtists.CreationDate) {
		return true
	}
	return false
}

func CheckFirstAlbum(allArtists *tools.Artists, Album1, Album2 string) bool {
	if len(Album1) == 0 && len(Album2) == 0 {
		return true
	}

	minYear, _ := strconv.Atoi(Album1)
	maxYear, _ := strconv.Atoi(Album2)
	firstAlbYear := strings.Split(allArtists.FirstAlbum, "-")[2]
	for i := minYear; i <= maxYear; i++ {
		if firstAlbYear == strconv.Itoa(i) {
			return true
		}
	}
	return false
}

func CheckNumberOfMembers(allArtists *tools.Artists, members []string) bool {
	if len(members) == 0 {
		return true
	}

	for _, e := range members {
		nb, _ := strconv.Atoi(e)
		if len(allArtists.Members) == nb {
			return true
		}
	}
	return false
}

func CheckLocations(locations *tools.Index, artists *tools.Artists, location string) bool {
	if location == "" {
		return true
	}
	if location == "seattle-usa" {
		location = "washington-usa"
	}
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
