package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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

	var allArtists *[]tools.Artists
	var locationsList []string
	url := "https://groupietrackers.herokuapp.com/api/artists"
	// get the api data
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("2")
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()
	// decode the jsone data
	err = json.NewDecoder(res.Body).Decode(&allArtists)
	if err != nil {
		fmt.Println("1")
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
		return
	}
	// Collect unique locations mn kol artist
	locationsSet := make(map[string]bool)

	for _, artist := range *allArtists {
		var locData *tools.LocationDataFilter
		err := helpers.Fetch_By_Id(artist.Locations, &locData)
		if err != nil {
			continue
		}

		for _, loc := range locData.Locations {
			cleaned := strings.TrimSpace(loc)
			if cleaned != "" {
				locationsSet[cleaned] = true
			}
		}
	}

	// Convert set to slice
	for loc := range locationsSet {
		locationsList = append(locationsList, loc)
	}
	// Prepare data
	finalData := struct {
		Artists   []tools.Artists
		Locations []string
	}{
		Artists:   *allArtists,
		Locations: locationsList,
	}

	helpers.RenderTemplates(w, "index.html", finalData, http.StatusOK)
}
