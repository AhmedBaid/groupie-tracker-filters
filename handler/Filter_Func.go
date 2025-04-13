package handler

import (
	"groupie/helpers"
	"groupie/tools"
	"net/http"
	"strconv"
)

func FilterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	minCreationDate := r.FormValue("minCreationDate")
	maxCreationDate := r.FormValue("maxCreationDate")
	firstAlbum1 := r.FormValue("firstAlbum1")
	firstAlbum2 := r.FormValue("firstAlbum2")
	numberOfMembers := r.Form["numberOfMembers"]
	locationsOfConcerts := r.FormValue("locationsOfConcerts")

	// fetch data from api
	var artistsData []tools.Artists
	err = helpers.Fetch_By_Id("https://groupietrackers.herokuapp.com/api/artists", &artistsData)
	if err != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
		return
	}

	var artists []tools.Artists
	data := tools.Data{}

	artistsFiltred(&artistsData, &artists, minCreationDate, maxCreationDate, firstAlbum1, firstAlbum2, locationsOfConcerts, numberOfMembers)

	Handle_data(&artistsData, &data)

	data.Artists = &artistsData

	helpers.RenderTemplates(w, "templates/index.html", data, http.StatusOK)
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
