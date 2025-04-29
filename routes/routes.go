package routes

import (
	"net/http"

	"groupie/handler"
)

func Route() {
	// our handlers
	http.HandleFunc("/", handler.Groupie_Func)
	http.HandleFunc("/static/", handler.Style_Func)
	http.HandleFunc("/details", handler.Detail_Card_Func)
	http.HandleFunc("/Filter", handler.FilterHandler)
}
