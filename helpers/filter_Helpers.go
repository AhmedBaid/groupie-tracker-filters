package helpers

import (
	"strconv"
	"strings"

	"groupie/tools"
)

func CheckCreationDate(artist *tools.Artists, minCr string, maxCr string) bool {
	minCrDate, _ := strconv.Atoi(minCr)
	maxCrDate, _ := strconv.Atoi(maxCr)
	if (minCrDate == 1987 && maxCrDate == 1987) || (artist.CreationDate >= minCrDate && maxCrDate >= artist.CreationDate) {
		return true
	}
	return false
}

func CheckFirstAlbum(artist *tools.Artists, Album1, Album2 string) bool {
	if len(Album1) == 0 && len(Album2) == 0 {
		return true
	}

	minYear, _ := strconv.Atoi(Album1)
	maxYear, _ := strconv.Atoi(Album2)
	firstAlbYear, _ := strconv.Atoi(strings.Split(artist.FirstAlbum, "-")[2])
	if firstAlbYear >= minYear && firstAlbYear <= maxYear {
		return true
	}

	return false
}

func CheckNumberOfMembers(artist *tools.Artists, members []string) bool {
	if len(members) == 0 {
		return true
	}

	for _, e := range members {
		nb, _ := strconv.Atoi(e)
		if len(artist.Members) == nb {
			return true
		}
	}
	return false
}

func CheckLocations(locations *tools.Index, artist *tools.Artists, location string) bool {
	location = strings.ToLower(strings.ReplaceAll(location, ", ", "-"))
	if location == "" {
		return true
	}
	if location == "seattle-usa" {
		location = "washington-usa"
	}
	for _, locations := range locations.Index {
		for _, loc := range locations.Locations {
			if loc == location {
				if locations.ID == artist.Id {
					return true
				}
			}
		}
	}
	return false
}

func ArtistsFiltred(allArtists *[]tools.Artists, minCrStr, maxCrStr, firstAlbumMin, firstAlbumMax, location string, members []string, locations tools.Index) *[]tools.Artists {
	filteredArtists := []tools.Artists{}
	var checkDate, checkFirstAlbum, checkMembers, checkLocations bool
	for _, artist := range *allArtists {
		checkDate = CheckCreationDate(&artist, minCrStr, maxCrStr)
		checkFirstAlbum = CheckFirstAlbum(&artist, firstAlbumMin, firstAlbumMax)
		checkMembers = CheckNumberOfMembers(&artist, members)
		checkLocations = CheckLocations(&locations, &artist, location)
		if checkMembers && checkFirstAlbum && checkDate && checkLocations {
			filteredArtists = append(filteredArtists, artist)
		}
	}
	return &filteredArtists
}
