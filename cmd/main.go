package main

import (
	"fmt"
	"net/http"

	"groupie/routes"
)

func main() {
	// Register handlers
	routes.Route()
	//  run   the server
	fmt.Println("Server running at http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server is already runing")
		return
	}
}
