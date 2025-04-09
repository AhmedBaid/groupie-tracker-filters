package helpers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func Fetch_By_Id(url string, target interface{}) error {
	// get the data from the url
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	// read the response body and decode it to the target variable
	err = json.NewDecoder(res.Body).Decode(target)
	if err != nil {
		return err
	}
	return nil
}

func ParseIntOrDefault(s string, def int) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		return def
	}
	return val
}

func ParseYear(date string) int {
	if len(date) < 4 {
		return 0
	}
	val, _ := strconv.Atoi(date[:4])
	return val
}

func ContainsLocation(locationStr string, target string) bool {
	target = strings.ToLower(target)
	locationStr = strings.ToLower(locationStr)
	return strings.Contains(locationStr, target)
}
