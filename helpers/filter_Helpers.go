package helpers

import (
	"strconv"
	"strings"

	"groupie/tools"
)

func GetCreattionDate(a *tools.Artists, min string, max string) bool {
	minV, _ := strconv.Atoi(min)
	maxV, _ := strconv.Atoi(max)
	if minV > maxV {
		minV, maxV = maxV, minV
	}
	if (minV == 1987 && maxV == 1987) || (a.CreationDate >= minV && maxV >= a.CreationDate) {
		return true
	}
	return false
}

func GetFirstAlbum(a *tools.Artists, y1, y2 string) bool {
	if len(y1) == 0 && len(y2) == 0 {
		return true
	}
	if y1 > y2 {
		y1, y2 = y2, y1
	}

	minyear, _ := strconv.Atoi(y1)
	maxyear, _ := strconv.Atoi(y2)
	for i := minyear; i <= maxyear; i++ {
		if strings.HasSuffix(a.FirstAlbum, "-"+strconv.Itoa(i)) {
			return true
		}
	}
	return false
}

func NumberOfMembers(a *tools.Artists, key []string) bool {
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

func LocationsOfConcert(l *tools.Index, a *tools.Artists, key string) bool {
	if key == "" {
		return true
	}
	if key == "seattle-usa" {
		key = "washington-usa"
	}
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
