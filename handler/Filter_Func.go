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
	fmt.Printf("firstAlbumMax: %v\n", firstAlbumMax)
	fmt.Printf("firstAlbumMin: %v\n", firstAlbumMin)
	fmt.Printf("numberOfMembers: %v\n", numberOfMembers)
	fmt.Printf("locationsOfConcerts: %v\n", locationsOfConcerts)

	// fetch data from api
	data := tools.Data{}
	var artistsData []tools.Artists
	var filtered []tools.Artists
	err = helpers.Fetch("https://groupietrackers.herokuapp.com/api/artists", &artistsData)
	if err != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
		return
	}

	Handle_data(&artistsData, &data)
	indexData := tools.Index{
		Index: []tools.Locations{},
	}

	artistsFiltred(&artistsData, &filtered, minCreationDate, maxCreationDate, firstAlbumMin, firstAlbumMax, locationsOfConcerts, numberOfMembers, indexData)
	data.Artists = &filtered
	fmt.Println("Number of filtered artists:", len(filtered))
	helpers.RenderTemplates(w, "index.html", data, http.StatusOK)
}

func artistsFiltred(all *[]tools.Artists, filtered *[]tools.Artists, minCrStr, maxCrStr, album1, album2, loc string, members []string, loci tools.Index) {
	hasDate := false
	hasFirstAlbum := false
	hasMembers := false
	hasLocations := false
	for _, artist := range *all {

		hasDate = helpers.GetCreattionDate(&artist, minCrStr, maxCrStr)
		hasFirstAlbum = helpers.GetFirstAlbum(&artist, album1, album2)
		hasMembers = helpers.NumberOfMembers(&artist, members)
		hasLocations = helpers.LocationsOfConcert(&loci, &artist, loc)
		if hasMembers && hasFirstAlbum && hasDate && hasLocations {
			*filtered = append(*filtered, artist)
		}
	}
}
