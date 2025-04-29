package handler

import (
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

	// fetch data from api
	data := tools.Data{}
	var artistsData *[]tools.Artists
	err = helpers.Fetch("https://groupietrackers.herokuapp.com/api/artists", &artistsData)
	if err != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
		return
	}

	Handle_data(artistsData, &data)
	filteredArtists := helpers.ArtistsFiltred(artistsData, minCreationDate, maxCreationDate, firstAlbumMin, firstAlbumMax, locationsOfConcerts, numberOfMembers, helpers.Locations)
	data.Artists = filteredArtists
	helpers.RenderTemplates(w, "filterPage.html", data, http.StatusOK)
}
