package handler

import (
	"fmt"
	"net/http"
	"strconv"

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
	var artistsData []tools.Artists
	err = helpers.Fetch("https://groupietrackers.herokuapp.com/api/artists", &artistsData)
	if err != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
		return
	}

	var artists []tools.Artists
	data := tools.Data{}

	artistsFiltred(&artistsData, &artists, minCreationDate, maxCreationDate, firstAlbumMin, firstAlbumMax, locationsOfConcerts, numberOfMembers)

	Handle_data(&artistsData, &data)

	data.Artists = &artistsData

	helpers.RenderTemplates(w, "index.html", data, http.StatusOK)
}

func artistsFiltred(all *[]tools.Artists, filtered *[]tools.Artists, minCrStr, maxCrStr, album1, album2, loc string, members []string) {
	minCr, _ := strconv.Atoi(minCrStr)
	maxCr, _ := strconv.Atoi(maxCrStr)
	minAlb, _ := strconv.Atoi(album1)
	maxAlb, _ := strconv.Atoi(album2)

	for _, artist := range *all {
		// filter Creation Date
		if minCr != 0 && artist.CreationDate < minCr {
			continue
		}
		if maxCr != 0 && artist.CreationDate > maxCr {
			continue
		}

		// filter First Album
		albumYear := helpers.ExtractYear(artist.FirstAlbum)
		if minAlb != 0 && albumYear < minAlb {
			continue
		}
		if maxAlb != 0 && albumYear > maxAlb {
			continue
		}

		// filter Members count
		if len(members) > 0 && !helpers.MemberMatch(len(artist.Members), members) {
			continue
		}

		// filter Location
		if loc != "" && !helpers.HasLocation(artist, loc) {
			continue
		}

		*filtered = append(*filtered, artist)
	}
}
