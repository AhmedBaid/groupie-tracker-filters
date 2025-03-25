package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"groupie/helpers"
	"groupie/tools"
)

// Groupie_Func handles the home page and applies filters if provided.
func Groupie_Func(w http.ResponseWriter, r *http.Request) {
	// Ensure the path is correct
	if r.URL.Path != "/" {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorNotFound, http.StatusNotFound)
		return
	}

	// Ensure method is GET
	if r.Method != http.MethodGet {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorMethodnotAll, http.StatusMethodNotAllowed)
		return
	}

	// Retrieve filter parameters from the request
	minCreationStr := r.URL.Query().Get("minCreationDate")
	maxCreationStr := r.URL.Query().Get("maxCreationDate")
	locationParam := r.URL.Query().Get("location")

	// Default filter values
	minCreation := 1958
	if minCreationStr != "" {
		if val, err := strconv.Atoi(minCreationStr); err == nil {
			minCreation = val
		}
	}

	maxCreation := 2015
	if maxCreationStr != "" {
		if val, err := strconv.Atoi(maxCreationStr); err == nil {
			maxCreation = val
		}
	}

	// Fetch all artists from API
	artistsURL := "https://groupietrackers.herokuapp.com/api/artists"
	res, err := http.Get(artistsURL)
	if err != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	var allArtists []tools.Artists
	if err = json.NewDecoder(res.Body).Decode(&allArtists); err != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
		return
	}

	// Filter artists based on query parameters
	var filteredArtists []tools.Artists
	for _, artist := range allArtists {
		if artist.CreationDate < minCreation || artist.CreationDate > maxCreation {
			continue
		}
		if locationParam != "" {
			if !strings.Contains(strings.ToLower(artist.Locations), strings.ToLower(locationParam)) {
				continue
			}
		}
		filteredArtists = append(filteredArtists, artist)
	}

	// Fetch available locations from API
	locationsURL := "https://groupietrackers.herokuapp.com/api/locations"
	resLoc, err := http.Get(locationsURL)
	if err != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
		return
	}
	defer resLoc.Body.Close()

	var locations []string
	if err = json.NewDecoder(resLoc.Body).Decode(&locations); err != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
		return
	}

	// Send filtered results to the home template
	data := tools.FilterArtists{
		Allartists:      filteredArtists,
		MinCreationDate: minCreation,
		MaxCreationDate: maxCreation,
		Locations:       locations,
	}

	helpers.RenderTemplates(w, "index.html", data, http.StatusOK)
}
