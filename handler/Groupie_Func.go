package handler

import (
	"net/http"

	"groupie/helpers"
	tools "groupie/tools"
)

func Groupie_Func(w http.ResponseWriter, r *http.Request) {
	// check the path
	if r.URL.Path != "/" {
		// execute the not found  template
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorNotFound, http.StatusNotFound)
		return
	}
	// check the methd
	if r.Method != http.MethodGet {
		// execute the not found  template
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorMethodnotAll, http.StatusMethodNotAllowed)
		return
	}
	finalData := &tools.Data{}

	// initialise the variables
	var allArtists *[]tools.Artists
	url := "https://groupietrackers.herokuapp.com/api/artists"
	err := helpers.Fetch(url, &allArtists)
	if err != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
		return
	}
	Handle_data(allArtists, finalData)
	finalData.Artists = allArtists
	helpers.RenderTemplates(w, "index.html", finalData, http.StatusOK)
}
