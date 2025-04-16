package handler

import (
	"fmt"
	"net/http"

	"groupie/helpers"
	"groupie/tools"
)

func FilterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorMethodnotAll, http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorBadReq, http.StatusBadRequest)
		return
	}

	minCreationDate := r.FormValue("Crmin")
	maxCreationDate := r.FormValue("Crmax")
	firstAlbumMin := r.FormValue("album-min")
	firstAlbumMax := r.FormValue("album-max")
	numberOfMembers := r.Form["members"]
	locationsOfConcerts := r.FormValue("location")
	fmt.Printf("minCreationDate: %v\n", minCreationDate)
	fmt.Printf("maxCreationDate: %v\n", maxCreationDate)


	// fetch data from api
	data := tools.Data{}
	var artistsData *[]tools.Artists
	err = helpers.Fetch("https://groupietrackers.herokuapp.com/api/artists", &artistsData)
	if err != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
		return
	}

	Handle_data(artistsData, &data)
	filteredArtists := artistsFiltred(artistsData, minCreationDate, maxCreationDate, firstAlbumMin, firstAlbumMax, locationsOfConcerts, numberOfMembers, helpers.Locations)
	data.Artists = filteredArtists
	helpers.RenderTemplates(w, "filterPage.html", data, http.StatusOK)
}

func artistsFiltred(allArtists *[]tools.Artists, minCrStr, maxCrStr, album1, album2, location string, members []string, locations tools.Index) *[]tools.Artists {
	filteredArtists := []tools.Artists{}
	for _, artist := range *allArtists {
		hasDate := helpers.CheckCreationDate(&artist, minCrStr, maxCrStr)
		hasFirstAlbum := helpers.CheckFirstAlbum(&artist, album1, album2)
		hasMembers := helpers.CheckNumberOfMembers(&artist, members)
		hasLocations := helpers.CheckLocations(&locations, &artist, location)
		if hasMembers && hasFirstAlbum && hasDate && hasLocations {
			filteredArtists = append(filteredArtists, artist)
		}
	}
	return &filteredArtists
}
