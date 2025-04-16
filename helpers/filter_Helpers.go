package helpers

import (
	"strconv"
	"strings"

	"groupie/tools"
)

func CheckCreationDate(a *tools.Artists, min string, max string) bool {
	minV, _ := strconv.Atoi(min)
	maxV, _ := strconv.Atoi(max)
	if (minV == 1987 && maxV == 1987) || (a.CreationDate >= minV && maxV >= a.CreationDate) {
		return true
	}
	return false
}

func CheckFirstAlbum(a *tools.Artists, y1, y2 string) bool {
	if len(y1) == 0 && len(y2) == 0 {
		return true
	}

	minyear, _ := strconv.Atoi(y1)
	maxyear, _ := strconv.Atoi(y2)
	firstAlbYear := strings.Split(a.FirstAlbum, "-")[2]
	for i := minyear; i <= maxyear; i++ {
		if firstAlbYear == strconv.Itoa(i) {
			return true
		}
	}
	return false
}

func CheckNumberOfMembers(a *tools.Artists, key []string) bool {
	if len(key) == 0 {
		return true
	}

	for _, e := range key {
		nb, _ := strconv.Atoi(e)
		if len(a.Members) == nb {
			return true
		}
	}
	return false
}

func CheckLocations(l *tools.Index, a *tools.Artists, key string) bool {
	if key == "" {
		return true
	}
	// if key == "seattle-usa" {
	// 	key = "washington-usa"
	// }
	for _, locations := range l.Index {
		for _, adress := range locations.Locations {
			if adress == key {
				if locations.ID == a.Id {
					return true
				}
			}
		}
	}
	return false
}
