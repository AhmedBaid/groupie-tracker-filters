package handler

import (
	"fmt"
	"net/http"
)

func Filter_Func(w http.ResponseWriter,r *http.Request) {
	fmt.Println("hello")
}
