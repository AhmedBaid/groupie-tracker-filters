package handler

import (
	"encoding/json"
	"net/http"

	"groupie/helpers"
	"groupie/tools"
)

func FilterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorMethodnotAll, http.StatusMethodNotAllowed)
		return
	}

	// Load all artists from API
	url := "https://groupietrackers.herokuapp.com/api/artists"
	res, err := http.Get(url)
	if err != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	var artists []tools.Artists
	err = json.NewDecoder(res.Body).Decode(&artists)
	if err != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
		return
	}

	// Get form values
	creationMin := helpers.ParseIntOrDefault(r.FormValue("creation-min"), 1958)
	creationMax := helpers.ParseIntOrDefault(r.FormValue("creation-max"), 2015)
	albumMin := helpers.ParseIntOrDefault(r.FormValue("album-min"), 1958)
	albumMax := helpers.ParseIntOrDefault(r.FormValue("album-max"), 2015)
	location := r.FormValue("locations")

	// Checkbox members: manually extract multiple values
	r.ParseForm() // 👈 this is important for FormValue to get multiple checkboxes
	memberValues := r.Form["members"]

	memberMap := make(map[int]bool)
	for _, m := range memberValues {
		val := helpers.ParseIntOrDefault(m, -1)
		if val > 0 {
			memberMap[val] = true
		}
	}

	// Filter logic
	var filtered []tools.Artists
	for _, artist := range artists {
		if artist.CreationDate < creationMin || artist.CreationDate > creationMax {
			continue
		}
		albumYear := helpers.ParseYear(artist.FirstAlbum)
		if albumYear < albumMin || albumYear > albumMax {
			continue
		}
		if len(memberMap) > 0 && !memberMap[len(artist.Members)] {
			continue
		}
		if location != "" && !helpers.ContainsLocation(artist.Locations, location) {
			continue
		}
		
		filtered = append(filtered, artist)
	}

	// Render filtered result
	helpers.RenderTemplates(w, "index.html", filtered, http.StatusOK)
}
